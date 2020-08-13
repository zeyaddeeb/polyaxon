// Copyright 2018-2020 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/**
 * Polyaxon SDKs and REST API specification.
 * Polyaxon SDKs and REST API specification.
 *
 * The version of the OpenAPI document: 1.1.7
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */


import ApiClient from "../ApiClient";
import RuntimeError from '../model/RuntimeError';
import V1Schemas from '../model/V1Schemas';

/**
* SchemasV1 service.
* @module api/SchemasV1Api
* @version 1.1.7
*/
export default class SchemasV1Api {

    /**
    * Constructs a new SchemasV1Api. 
    * Polyaxon sdk
    * @alias module:api/SchemasV1Api
    * @class
    * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
    * default to {@link module:ApiClient#instance} if unspecified.
    */
    constructor(apiClient) {
        this.apiClient = apiClient || ApiClient.instance;
    }


    /**
     * Callback function to receive the result of the noOp operation.
     * @callback module:api/SchemasV1Api~noOpCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1Schemas} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Get random
     * @param {module:api/SchemasV1Api~noOpCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1Schemas}
     */
    noOp(callback) {
      let postBody = null;

      let pathParams = {
      };
      let queryParams = {
      };
      let headerParams = {
      };
      let formParams = {
      };

      let authNames = ['ApiKey'];
      let contentTypes = [];
      let accepts = ['application/json'];
      let returnType = V1Schemas;
      return this.apiClient.callApi(
        '/schemas', 'GET',
        pathParams, queryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, null, callback
      );
    }


}
