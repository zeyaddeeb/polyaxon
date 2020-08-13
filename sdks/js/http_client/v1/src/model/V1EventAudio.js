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

/**
 * The V1EventAudio model module.
 * @module model/V1EventAudio
 * @version 1.1.7
 */
class V1EventAudio {
    /**
     * Constructs a new <code>V1EventAudio</code>.
     * @alias module:model/V1EventAudio
     */
    constructor() { 
        
        V1EventAudio.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>V1EventAudio</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1EventAudio} obj Optional instance to populate.
     * @return {module:model/V1EventAudio} The populated <code>V1EventAudio</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1EventAudio();

            if (data.hasOwnProperty('sample_rate')) {
                obj['sample_rate'] = ApiClient.convertToType(data['sample_rate'], 'Number');
            }
            if (data.hasOwnProperty('num_channels')) {
                obj['num_channels'] = ApiClient.convertToType(data['num_channels'], 'Number');
            }
            if (data.hasOwnProperty('length_frames')) {
                obj['length_frames'] = ApiClient.convertToType(data['length_frames'], 'Number');
            }
            if (data.hasOwnProperty('content_type')) {
                obj['content_type'] = ApiClient.convertToType(data['content_type'], 'String');
            }
            if (data.hasOwnProperty('path')) {
                obj['path'] = ApiClient.convertToType(data['path'], 'String');
            }
        }
        return obj;
    }


}

/**
 * Sample rate of the audio in Hz.
 * @member {Number} sample_rate
 */
V1EventAudio.prototype['sample_rate'] = undefined;

/**
 * Number of channels of audio.
 * @member {Number} num_channels
 */
V1EventAudio.prototype['num_channels'] = undefined;

/**
 * Length of the audio in frames (samples per channel).
 * @member {Number} length_frames
 */
V1EventAudio.prototype['length_frames'] = undefined;

/**
 * @member {String} content_type
 */
V1EventAudio.prototype['content_type'] = undefined;

/**
 * @member {String} path
 */
V1EventAudio.prototype['path'] = undefined;






export default V1EventAudio;

