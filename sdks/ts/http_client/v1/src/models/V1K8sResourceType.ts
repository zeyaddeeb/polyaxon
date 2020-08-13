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
    V1K8sResourceSchema,
    V1K8sResourceSchemaFromJSON,
    V1K8sResourceSchemaFromJSONTyped,
    V1K8sResourceSchemaToJSON,
} from './';

/**
 * 
 * @export
 * @interface V1K8sResourceType
 */
export interface V1K8sResourceType {
    /**
     * 
     * @type {string}
     * @memberof V1K8sResourceType
     */
    name?: string;
    /**
     * 
     * @type {V1K8sResourceSchema}
     * @memberof V1K8sResourceType
     */
    schema?: V1K8sResourceSchema;
    /**
     * 
     * @type {boolean}
     * @memberof V1K8sResourceType
     */
    is_requested?: boolean;
}

export function V1K8sResourceTypeFromJSON(json: any): V1K8sResourceType {
    return V1K8sResourceTypeFromJSONTyped(json, false);
}

export function V1K8sResourceTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1K8sResourceType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'name': !exists(json, 'name') ? undefined : json['name'],
        'schema': !exists(json, 'schema') ? undefined : V1K8sResourceSchemaFromJSON(json['schema']),
        'is_requested': !exists(json, 'is_requested') ? undefined : json['is_requested'],
    };
}

export function V1K8sResourceTypeToJSON(value?: V1K8sResourceType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'name': value.name,
        'schema': V1K8sResourceSchemaToJSON(value.schema),
        'is_requested': value.is_requested,
    };
}


