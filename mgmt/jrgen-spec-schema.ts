// Generated with these commands:
// wget https://github.com/mzernetsch/jrgen/raw/v2.0.2/jrgen-spec.schema.json
// npx -p json-schema-to-typescript json2ts -i jrgen-spec.schema.json -o mgmt/jrgen-spec-schema.ts
// rm jrgen-spec.schema.json

/**
 * This file was automatically generated by json-schema-to-typescript.
 * DO NOT MODIFY IT BY HAND. Instead, modify the source JSONSchema file,
 * and run json-schema-to-typescript to regenerate this file.
 */

export interface JrgenSpecSchema {
  /**
   * Version of the jrgen schema.
   */
  jrgen: "1.0" | "1.1";
  /**
   * Version of the json-rpc protocol.
   */
  jsonrpc: "2.0";
  /**
   * Meta information about the api.
   */
  info: {
    /**
     * Name of the api.
     */
    title: string;
    /**
     * Description or usage information about the api.
     */
    description?: string | string[];
    /**
     * Current version of the api.
     */
    version: string;
    [k: string]: any;
  };
  /**
   * Global definitions for use in the api.
   */
  definitions?: {
    [k: string]: any;
  };
  /**
   * Definitions of the available procedures in the api. A key equals to the name of a procedure.
   */
  methods: {
    /**
     * Definition of an api procedure.
     *
     * This interface was referenced by `undefined`'s JSON-Schema definition
     * via the `patternProperty` "^.*$".
     */
    [k: string]: {
      /**
       * Short summary of what the procedure does.
       */
      summary: string;
      /**
       * Longer description of what the procedure does.
       */
      description?: string | string[];
      /**
       * Tags for grouping similar procedures.
       */
      tags?: string[];
      params?: any;
      result?: any;
      /**
       * Definition of possible error responses.
       */
      errors?: {
        /**
         * Description of what went wrong.
         */
        description?: string;
        /**
         * Unique error code.
         */
        code: number;
        /**
         * Unique error message,
         */
        message: string;
        data?: any;
        [k: string]: any;
      }[];
      [k: string]: any;
    };
  };
  [k: string]: any;
}
