package ping

import (
	"github.com/usnistgov/ndn-dpdk/app/fetch"
	"github.com/usnistgov/ndn-dpdk/app/pingclient"
	"github.com/usnistgov/ndn-dpdk/app/pingserver"
	"github.com/usnistgov/ndn-dpdk/dpdk/ealthread"
	"github.com/usnistgov/ndn-dpdk/iface"
)

type Task struct {
	Face   iface.Face
	Server []*pingserver.Server
	Client *pingclient.Client
	Fetch  *fetch.Fetcher
}

func newTask(face iface.Face, cfg TaskConfig) (task Task, e error) {
	socket := face.NumaSocket()
	task.Face = face

	if cfg.Server != nil {
		nThreads := cfg.Server.NThreads
		if nThreads <= 0 {
			nThreads = 1
		}
		for i := 0; i < nThreads; i++ {
			if server, e := pingserver.New(task.Face, i, cfg.Server.Config); e != nil {
				return Task{}, e
			} else {
				server.SetLCore(ealthread.DefaultAllocator.Alloc(LCoreRole_Server, socket))
				task.Server = append(task.Server, server)
			}
		}
	}

	if cfg.Client != nil {
		if task.Client, e = pingclient.New(task.Face, *cfg.Client); e != nil {
			return Task{}, e
		}
		task.Client.SetLCores(ealthread.DefaultAllocator.Alloc(LCoreRole_ClientRx, socket), ealthread.DefaultAllocator.Alloc(LCoreRole_ClientTx, socket))
	} else if cfg.Fetch != nil {
		if task.Fetch, e = fetch.New(task.Face, *cfg.Fetch); e != nil {
			return Task{}, e
		}
		for i, last := 0, task.Fetch.CountThreads(); i < last; i++ {
			task.Fetch.Thread(i).SetLCore(ealthread.DefaultAllocator.Alloc(LCoreRole_ClientRx, socket))
		}
	}

	return task, nil
}

func (task *Task) configureDemux(demuxI, demuxD, demuxN *iface.InputDemux) {
	if nServers := len(task.Server); nServers > 0 {
		demuxI.InitRoundrobin(nServers)
		for i, server := range task.Server {
			demuxI.SetDest(i, server.RxQueue())
		}
	}

	if task.Client != nil {
		demuxD.InitFirst()
		demuxN.InitFirst()
		q := task.Client.RxQueue()
		demuxD.SetDest(0, q)
		demuxN.SetDest(0, q)
	} else if task.Fetch != nil {
		demuxD.InitToken()
		demuxN.InitToken()
		for i, last := 0, task.Fetch.CountProcs(); i < last; i++ {
			q := task.Fetch.RxQueue(i)
			demuxD.SetDest(i, q)
			demuxN.SetDest(i, q)
		}
	}
}

func (task *Task) Launch() {
	for _, server := range task.Server {
		server.Launch()
	}
	if task.Client != nil {
		task.Client.Launch()
	}
}

func (task *Task) Close() error {
	for _, server := range task.Server {
		server.Close()
	}
	if task.Client != nil {
		task.Client.Close()
	}
	if task.Fetch != nil {
		task.Fetch.Close()
	}
	task.Face.Close()
	return nil
}
