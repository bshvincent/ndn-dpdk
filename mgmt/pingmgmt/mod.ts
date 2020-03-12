import { Counters as FetchCounters_ } from "../../app/fetch/mod.js";
import { Counters as ClientCounters_ } from "../../app/pingclient/mod.js";
import { Index } from "../../core/mod.js";
import { Milliseconds, Nanoseconds } from "../../core/nnduration/mod.js";
import { Name } from "../../ndn/mod.js";

export interface IndexArg {
  Index: Index;
}

export interface ClientStartArgs extends IndexArg {
  Interval: Nanoseconds;

  /**
   * @default false
   */
  ClearCounters?: boolean;
}

export interface ClientStopArgs extends IndexArg {
  RxDelay?: Nanoseconds;
}

export type ClientCounters = ClientCounters_;

export interface PingClientMgmt {
  List: {args: {}, reply: Index[]};
  Start: {args: ClientStartArgs, reply: {}};
  Stop: {args: ClientStopArgs, reply: {}};
  ReadCounters: {args: IndexArg, reply: ClientCounters};
}

export interface FetchIndexArg extends IndexArg {
  FetchId: Index;
}

export type FetchCounters = FetchCounters_;

export interface FetchTemplate {
  Prefix: Name;
  CanBePrefix?: boolean;
}

export interface FetchBenchmarkArgs extends FetchIndexArg {
  Templates: FetchTemplate[];
  Warmup: Milliseconds;
  Interval: Milliseconds;
  Count: number;
}

export interface FetchBenchmarkReply {
  Counters: FetchCounters[];
  Goodput: number;
}

export interface FetchMgmt {
  List: {args: {}, reply: FetchIndexArg[]};
  Benchmark: {args: FetchBenchmarkArgs, reply: FetchBenchmarkReply};
  ReadCounters: {args: FetchIndexArg, reply: FetchCounters};
}