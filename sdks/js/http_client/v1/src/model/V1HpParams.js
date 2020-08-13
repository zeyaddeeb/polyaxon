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

import ApiClient from '../ApiClient';
import V1HpChoice from './V1HpChoice';
import V1HpGeomSpace from './V1HpGeomSpace';
import V1HpLinSpace from './V1HpLinSpace';
import V1HpLogNormal from './V1HpLogNormal';
import V1HpLogSpace from './V1HpLogSpace';
import V1HpLogUniform from './V1HpLogUniform';
import V1HpNormal from './V1HpNormal';
import V1HpPChoice from './V1HpPChoice';
import V1HpQLogNormal from './V1HpQLogNormal';
import V1HpQLogUniform from './V1HpQLogUniform';
import V1HpQNormal from './V1HpQNormal';
import V1HpQUniform from './V1HpQUniform';
import V1HpRange from './V1HpRange';
import V1HpUniform from './V1HpUniform';

/**
 * The V1HpParams model module.
 * @module model/V1HpParams
 * @version 1.1.7
 */
class V1HpParams {
    /**
     * Constructs a new <code>V1HpParams</code>.
     * @alias module:model/V1HpParams
     */
    constructor() { 
        
        V1HpParams.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>V1HpParams</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1HpParams} obj Optional instance to populate.
     * @return {module:model/V1HpParams} The populated <code>V1HpParams</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1HpParams();

            if (data.hasOwnProperty('choice')) {
                obj['choice'] = V1HpChoice.constructFromObject(data['choice']);
            }
            if (data.hasOwnProperty('pchoice')) {
                obj['pchoice'] = V1HpPChoice.constructFromObject(data['pchoice']);
            }
            if (data.hasOwnProperty('range')) {
                obj['range'] = V1HpRange.constructFromObject(data['range']);
            }
            if (data.hasOwnProperty('linspace')) {
                obj['linspace'] = V1HpLinSpace.constructFromObject(data['linspace']);
            }
            if (data.hasOwnProperty('logspace')) {
                obj['logspace'] = V1HpLogSpace.constructFromObject(data['logspace']);
            }
            if (data.hasOwnProperty('geomspace')) {
                obj['geomspace'] = V1HpGeomSpace.constructFromObject(data['geomspace']);
            }
            if (data.hasOwnProperty('uniform')) {
                obj['uniform'] = V1HpUniform.constructFromObject(data['uniform']);
            }
            if (data.hasOwnProperty('quniform')) {
                obj['quniform'] = V1HpQUniform.constructFromObject(data['quniform']);
            }
            if (data.hasOwnProperty('loguniform')) {
                obj['loguniform'] = V1HpLogUniform.constructFromObject(data['loguniform']);
            }
            if (data.hasOwnProperty('qloguniform')) {
                obj['qloguniform'] = V1HpQLogUniform.constructFromObject(data['qloguniform']);
            }
            if (data.hasOwnProperty('normal')) {
                obj['normal'] = V1HpNormal.constructFromObject(data['normal']);
            }
            if (data.hasOwnProperty('qnormal')) {
                obj['qnormal'] = V1HpQNormal.constructFromObject(data['qnormal']);
            }
            if (data.hasOwnProperty('lognormal')) {
                obj['lognormal'] = V1HpLogNormal.constructFromObject(data['lognormal']);
            }
            if (data.hasOwnProperty('qlognormal')) {
                obj['qlognormal'] = V1HpQLogNormal.constructFromObject(data['qlognormal']);
            }
        }
        return obj;
    }


}

/**
 * @member {module:model/V1HpChoice} choice
 */
V1HpParams.prototype['choice'] = undefined;

/**
 * @member {module:model/V1HpPChoice} pchoice
 */
V1HpParams.prototype['pchoice'] = undefined;

/**
 * @member {module:model/V1HpRange} range
 */
V1HpParams.prototype['range'] = undefined;

/**
 * @member {module:model/V1HpLinSpace} linspace
 */
V1HpParams.prototype['linspace'] = undefined;

/**
 * @member {module:model/V1HpLogSpace} logspace
 */
V1HpParams.prototype['logspace'] = undefined;

/**
 * @member {module:model/V1HpGeomSpace} geomspace
 */
V1HpParams.prototype['geomspace'] = undefined;

/**
 * @member {module:model/V1HpUniform} uniform
 */
V1HpParams.prototype['uniform'] = undefined;

/**
 * @member {module:model/V1HpQUniform} quniform
 */
V1HpParams.prototype['quniform'] = undefined;

/**
 * @member {module:model/V1HpLogUniform} loguniform
 */
V1HpParams.prototype['loguniform'] = undefined;

/**
 * @member {module:model/V1HpQLogUniform} qloguniform
 */
V1HpParams.prototype['qloguniform'] = undefined;

/**
 * @member {module:model/V1HpNormal} normal
 */
V1HpParams.prototype['normal'] = undefined;

/**
 * @member {module:model/V1HpQNormal} qnormal
 */
V1HpParams.prototype['qnormal'] = undefined;

/**
 * @member {module:model/V1HpLogNormal} lognormal
 */
V1HpParams.prototype['lognormal'] = undefined;

/**
 * @member {module:model/V1HpQLogNormal} qlognormal
 */
V1HpParams.prototype['qlognormal'] = undefined;






export default V1HpParams;

