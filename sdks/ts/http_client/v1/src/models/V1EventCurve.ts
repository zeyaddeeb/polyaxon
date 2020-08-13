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

/* tslint:disable */
/* eslint-disable */
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
 */

import { exists, mapValues } from '../runtime';
import {
    V1EventCurveKind,
    V1EventCurveKindFromJSON,
    V1EventCurveKindFromJSONTyped,
    V1EventCurveKindToJSON,
} from './';

/**
 * 
 * @export
 * @interface V1EventCurve
 */
export interface V1EventCurve {
    /**
     * 
     * @type {V1EventCurveKind}
     * @memberof V1EventCurve
     */
    kind?: V1EventCurveKind;
    /**
     * 
     * @type {Array<number>}
     * @memberof V1EventCurve
     */
    x?: Array<number>;
    /**
     * 
     * @type {Array<number>}
     * @memberof V1EventCurve
     */
    y?: Array<number>;
    /**
     * 
     * @type {string}
     * @memberof V1EventCurve
     */
    annotation?: string;
}

export function V1EventCurveFromJSON(json: any): V1EventCurve {
    return V1EventCurveFromJSONTyped(json, false);
}

export function V1EventCurveFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1EventCurve {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'kind': !exists(json, 'kind') ? undefined : V1EventCurveKindFromJSON(json['kind']),
        'x': !exists(json, 'x') ? undefined : json['x'],
        'y': !exists(json, 'y') ? undefined : json['y'],
        'annotation': !exists(json, 'annotation') ? undefined : json['annotation'],
    };
}

export function V1EventCurveToJSON(value?: V1EventCurve | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'kind': V1EventCurveKindToJSON(value.kind),
        'x': value.x,
        'y': value.y,
        'annotation': value.annotation,
    };
}


