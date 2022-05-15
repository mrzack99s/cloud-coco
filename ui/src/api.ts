/* eslint-disable */
/* tslint:disable */
/*
 * ---------------------------------------------------------------
 * ## THIS FILE WAS GENERATED VIA SWAGGER-TYPESCRIPT-API        ##
 * ##                                                           ##
 * ## AUTHOR: acacode                                           ##
 * ## SOURCE: https://github.com/acacode/swagger-typescript-api ##
 * ---------------------------------------------------------------
 */

export interface AuthenticationCredentialParams {
  email?: string;
  password?: string;
}

export interface AuthenticationTokenResponse {
  access_token?: string;
  access_token_expires?: string;
  need_chpasswd?: boolean;
}

export type GinH = Record<string, any>;

export interface GormDeletedAt {
  time?: string;

  /** Valid is true if Time is not NULL */
  valid?: boolean;
}

export interface ModelsDirectories {
  created_at?: string;
  deleted_at?: GormDeletedAt;
  disabled?: boolean;
  id?: number;
  name?: string;
  subscriptions?: ModelsSubscriptions[];
  updated_at?: string;
  users?: ModelsDirectoriesUsers[];
  uuid?: string;
}

export interface ModelsDirectoriesUsers {
  directory?: ModelsDirectories;
  directory_id?: number;
  user_id?: number;
  users?: ModelsUsers;
}

export interface ModelsPermissions {
  created_at?: string;
  deleted_at?: GormDeletedAt;
  id?: number;
  name?: string;
  permissions?: ModelsRolesPermissions[];
  updated_at?: string;
  uuid?: string;
}

export interface ModelsRBACResourcePools {
  resource_pool?: ModelsResourcePools;
  resource_pool_id?: number;
  role?: ModelsRoles;
  role_id?: number;
  user?: ModelsUsers;
  user_id?: number;
}

export interface ModelsRBACSubscriptions {
  role?: ModelsRoles;
  role_id?: number;
  subscription?: ModelsSubscriptions;
  subscription_id?: number;
  user?: ModelsUsers;
  user_id?: number;
}

export interface ModelsResourcePools {
  created_at?: string;
  deleted_at?: GormDeletedAt;
  id?: number;
  name?: string;
  rbac?: ModelsRBACResourcePools[];
  resources?: ModelsResources[];
  subscription?: ModelsSubscriptions;
  subscription_id?: string;
  updated_at?: string;
  uuid?: string;
}

export interface ModelsResources {
  created_at?: string;
  deleted_at?: GormDeletedAt;
  id?: number;
  name?: string;
  resource_pool?: ModelsResourcePools;
  resource_pool_id?: number;
  resource_status?: ModelsResourcesStatus;
  resource_status_id?: number;
  service?: ModelsServiceVersions;
  service_id?: string;
  updated_at?: string;
  uuid?: string;
}

export interface ModelsResourcesStatus {
  created_at?: string;
  deleted_at?: GormDeletedAt;
  id?: number;
  name?: string;
  resources?: ModelsResources[];
  updated_at?: string;
  uuid?: string;
}

export interface ModelsRoles {
  created_at?: string;
  deleted_at?: GormDeletedAt;
  id?: number;
  name?: string;
  permissions?: ModelsRolesPermissions[];
  rbac_resource_pools?: ModelsRBACResourcePools[];
  rbac_subscriptions?: ModelsRBACSubscriptions[];
  updated_at?: string;
  uuid?: string;
}

export interface ModelsRolesPermissions {
  permission?: ModelsPermissions;
  permission_id?: number;
  role?: ModelsRoles;
  role_id?: number;
}

export interface ModelsServiceVersions {
  created_at?: string;
  deleted_at?: GormDeletedAt;
  id?: number;
  image?: string;
  resources?: ModelsResources[];
  service?: ModelsServices;
  service_id?: number;
  updated_at?: string;
  uuid?: string;
  version?: string;
}

export interface ModelsServices {
  created_at?: string;
  deleted_at?: GormDeletedAt;
  id?: number;
  name?: string;
  updated_at?: string;
  uuid?: string;
  versions?: ModelsServiceVersions[];
}

export interface ModelsSubscriptions {
  created_at?: string;
  deleted_at?: GormDeletedAt;
  directory?: ModelsDirectories;
  directory_id?: number;
  disabled?: boolean;
  id?: number;
  name?: string;
  rbac?: ModelsRBACSubscriptions[];
  resource_pools?: ModelsResourcePools[];
  updated_at?: string;
  uuid?: string;
}

export interface ModelsUsers {
  created_at?: string;
  deleted_at?: GormDeletedAt;
  directories?: ModelsDirectoriesUsers[];
  email?: string;
  first_name?: string;
  id?: number;
  last_name?: string;
  need_chpasswd?: boolean;
  passwd_checksum?: string;
  rbac_resource_pools?: ModelsRBACResourcePools[];
  rbac_subscriptions?: ModelsRBACSubscriptions[];
  temporary_password?: string;
  updated_at?: string;
  uuid?: string;
}

export interface TypesArrayResponse {
  record_count?: number;
  record_list?: any;
}

export interface TypesSetupParams {
  email?: string;
  first_name?: string;
  last_name?: string;
  password?: string;
}

export interface TypesUserChangePasswdParams {
  new_password?: string;
  old_password?: string;
  uuid?: number;
}

export interface TypesUserResetPasswdParams {
  uuid?: string;
}

export type QueryParamsType = Record<string | number, any>;
export type ResponseFormat = keyof Omit<Body, "body" | "bodyUsed">;

export interface FullRequestParams extends Omit<RequestInit, "body"> {
  /** set parameter to `true` for call `securityWorker` for this request */
  secure?: boolean;
  /** request path */
  path: string;
  /** content type of request body */
  type?: ContentType;
  /** query params */
  query?: QueryParamsType;
  /** format of response (i.e. response.json() -> format: "json") */
  format?: ResponseFormat;
  /** request body */
  body?: unknown;
  /** base url */
  baseUrl?: string;
  /** request cancellation token */
  cancelToken?: CancelToken;
}

export type RequestParams = Omit<FullRequestParams, "body" | "method" | "query" | "path">;

export interface ApiConfig<SecurityDataType = unknown> {
  baseUrl?: string;
  baseApiParams?: Omit<RequestParams, "baseUrl" | "cancelToken" | "signal">;
  securityWorker?: (securityData: SecurityDataType | null) => Promise<RequestParams | void> | RequestParams | void;
  customFetch?: typeof fetch;
}

export interface HttpResponse<D extends unknown, E extends unknown = unknown> extends Response {
  data: D;
  error: E;
}

type CancelToken = Symbol | string | number;

export enum ContentType {
  Json = "application/json",
  FormData = "multipart/form-data",
  UrlEncoded = "application/x-www-form-urlencoded",
}

export class HttpClient<SecurityDataType = unknown> {
  public baseUrl: string = "/api";
  private securityData: SecurityDataType | null = null;
  private securityWorker?: ApiConfig<SecurityDataType>["securityWorker"];
  private abortControllers = new Map<CancelToken, AbortController>();
  private customFetch = (...fetchParams: Parameters<typeof fetch>) => fetch(...fetchParams);

  private baseApiParams: RequestParams = {
    credentials: "same-origin",
    headers: {},
    redirect: "follow",
    referrerPolicy: "no-referrer",
  };

  constructor(apiConfig: ApiConfig<SecurityDataType> = {}) {
    Object.assign(this, apiConfig);
  }

  public setSecurityData = (data: SecurityDataType | null) => {
    this.securityData = data;
  };

  private encodeQueryParam(key: string, value: any) {
    const encodedKey = encodeURIComponent(key);
    return `${encodedKey}=${encodeURIComponent(typeof value === "number" ? value : `${value}`)}`;
  }

  private addQueryParam(query: QueryParamsType, key: string) {
    return this.encodeQueryParam(key, query[key]);
  }

  private addArrayQueryParam(query: QueryParamsType, key: string) {
    const value = query[key];
    return value.map((v: any) => this.encodeQueryParam(key, v)).join("&");
  }

  protected toQueryString(rawQuery?: QueryParamsType): string {
    const query = rawQuery || {};
    const keys = Object.keys(query).filter((key) => "undefined" !== typeof query[key]);
    return keys
      .map((key) => (Array.isArray(query[key]) ? this.addArrayQueryParam(query, key) : this.addQueryParam(query, key)))
      .join("&");
  }

  protected addQueryParams(rawQuery?: QueryParamsType): string {
    const queryString = this.toQueryString(rawQuery);
    return queryString ? `?${queryString}` : "";
  }

  private contentFormatters: Record<ContentType, (input: any) => any> = {
    [ContentType.Json]: (input: any) =>
      input !== null && (typeof input === "object" || typeof input === "string") ? JSON.stringify(input) : input,
    [ContentType.FormData]: (input: any) =>
      Object.keys(input || {}).reduce((formData, key) => {
        const property = input[key];
        formData.append(
          key,
          property instanceof Blob
            ? property
            : typeof property === "object" && property !== null
            ? JSON.stringify(property)
            : `${property}`,
        );
        return formData;
      }, new FormData()),
    [ContentType.UrlEncoded]: (input: any) => this.toQueryString(input),
  };

  private mergeRequestParams(params1: RequestParams, params2?: RequestParams): RequestParams {
    return {
      ...this.baseApiParams,
      ...params1,
      ...(params2 || {}),
      headers: {
        ...(this.baseApiParams.headers || {}),
        ...(params1.headers || {}),
        ...((params2 && params2.headers) || {}),
      },
    };
  }

  private createAbortSignal = (cancelToken: CancelToken): AbortSignal | undefined => {
    if (this.abortControllers.has(cancelToken)) {
      const abortController = this.abortControllers.get(cancelToken);
      if (abortController) {
        return abortController.signal;
      }
      return void 0;
    }

    const abortController = new AbortController();
    this.abortControllers.set(cancelToken, abortController);
    return abortController.signal;
  };

  public abortRequest = (cancelToken: CancelToken) => {
    const abortController = this.abortControllers.get(cancelToken);

    if (abortController) {
      abortController.abort();
      this.abortControllers.delete(cancelToken);
    }
  };

  public request = async <T = any, E = any>({
    body,
    secure,
    path,
    type,
    query,
    format,
    baseUrl,
    cancelToken,
    ...params
  }: FullRequestParams): Promise<HttpResponse<T, E>> => {
    const secureParams =
      ((typeof secure === "boolean" ? secure : this.baseApiParams.secure) &&
        this.securityWorker &&
        (await this.securityWorker(this.securityData))) ||
      {};
    const requestParams = this.mergeRequestParams(params, secureParams);
    const queryString = query && this.toQueryString(query);
    const payloadFormatter = this.contentFormatters[type || ContentType.Json];
    const responseFormat = format || requestParams.format;

    return this.customFetch(`${baseUrl || this.baseUrl || ""}${path}${queryString ? `?${queryString}` : ""}`, {
      ...requestParams,
      headers: {
        ...(type && type !== ContentType.FormData ? { "Content-Type": type } : {}),
        ...(requestParams.headers || {}),
      },
      signal: cancelToken ? this.createAbortSignal(cancelToken) : void 0,
      body: typeof body === "undefined" || body === null ? null : payloadFormatter(body),
    }).then(async (response) => {
      const r = response as HttpResponse<T, E>;
      r.data = null as unknown as T;
      r.error = null as unknown as E;

      const data = !responseFormat
        ? r
        : await response[responseFormat]()
            .then((data) => {
              if (r.ok) {
                r.data = data;
              } else {
                r.error = data;
              }
              return r;
            })
            .catch((e) => {
              r.error = e;
              return r;
            });

      if (cancelToken) {
        this.abortControllers.delete(cancelToken);
      }

      if (!response.ok) throw data;
      return data;
    });
  };
}

/**
 * @title Cloud COCO
 * @version 0.1.0
 * @license Apache License Version 2.0 (https://github.com/mrzack99s/cloud-coco)
 * @baseUrl /api
 * @contact
 *
 * This is a Cloud COCO API
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  authentication = {
    /**
     * @description Check credential for get access
     *
     * @tags Authentication
     * @name CheckCredential
     * @summary Check credential
     * @request POST:/authentication/check-credential
     */
    checkCredential: (params: AuthenticationCredentialParams, requestParams: RequestParams = {}) =>
      this.request<AuthenticationTokenResponse, GinH>({
        path: `/authentication/check-credential`,
        method: "POST",
        body: params,
        type: ContentType.Json,
        format: "json",
        ...requestParams,
      }),

    /**
     * @description Revoke token
     *
     * @tags Authentication
     * @name GetCredential
     * @summary Revoke token
     * @request GET:/authentication/get-credential
     * @secure
     */
    getCredential: (params: RequestParams = {}) =>
      this.request<ModelsUsers, GinH>({
        path: `/authentication/get-credential`,
        method: "GET",
        secure: true,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * @description Revoke token
     *
     * @tags Authentication
     * @name RevokeToken
     * @summary Revoke token
     * @request GET:/authentication/revoke-token
     * @secure
     */
    revokeToken: (params: RequestParams = {}) =>
      this.request<string, GinH>({
        path: `/authentication/revoke-token`,
        method: "GET",
        secure: true,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),
  };
  directories = {
    /**
     * @description Create Directories
     *
     * @tags Directories
     * @name CreateDirectories
     * @summary Create Directories
     * @request POST:/directories/create
     * @secure
     */
    createDirectories: (params: ModelsDirectories, requestParams: RequestParams = {}) =>
      this.request<ModelsDirectories, GinH>({
        path: `/directories/create`,
        method: "POST",
        body: params,
        secure: true,
        type: ContentType.Json,
        format: "json",
        ...requestParams,
      }),

    /**
     * @description Get Directories by offset
     *
     * @tags Directories
     * @name GetDirectoriesByOffset
     * @summary Get Directories by offset
     * @request GET:/directories/get-by-offset
     * @secure
     */
    getDirectoriesByOffset: (query: { offset: number; limit: number }, params: RequestParams = {}) =>
      this.request<TypesArrayResponse, GinH>({
        path: `/directories/get-by-offset`,
        method: "GET",
        query: query,
        secure: true,
        format: "json",
        ...params,
      }),

    /**
     * @description Get Directories by uuid
     *
     * @tags Directories
     * @name GetDirectoriesByUuid
     * @summary Get Directories by uuid
     * @request GET:/directories/get/{uuid}
     * @secure
     */
    getDirectoriesByUuid: (uuid: string, params: RequestParams = {}) =>
      this.request<ModelsDirectories, GinH>({
        path: `/directories/get/${uuid}`,
        method: "GET",
        secure: true,
        format: "json",
        ...params,
      }),

    /**
     * @description Hard Delete Directories
     *
     * @tags Directories
     * @name HardDeleteDirectories
     * @summary Hard Delete Directories
     * @request DELETE:/directories/hard-delete/{uuid}
     * @secure
     */
    hardDeleteDirectories: (uuid: string, params: RequestParams = {}) =>
      this.request<string, GinH>({
        path: `/directories/hard-delete/${uuid}`,
        method: "DELETE",
        secure: true,
        format: "json",
        ...params,
      }),

    /**
     * @description Soft Delete Directories
     *
     * @tags Directories
     * @name SoftDeleteDirectories
     * @summary Soft Delete Directories
     * @request DELETE:/directories/soft-delete/{uuid}
     * @secure
     */
    softDeleteDirectories: (uuid: string, params: RequestParams = {}) =>
      this.request<string, GinH>({
        path: `/directories/soft-delete/${uuid}`,
        method: "DELETE",
        secure: true,
        format: "json",
        ...params,
      }),

    /**
     * @description Update Directories
     *
     * @tags Directories
     * @name UpdateDirectories
     * @summary Update Directories
     * @request POST:/directories/update
     * @secure
     */
    updateDirectories: (params: ModelsDirectories, requestParams: RequestParams = {}) =>
      this.request<ModelsDirectories, GinH>({
        path: `/directories/update`,
        method: "POST",
        body: params,
        secure: true,
        type: ContentType.Json,
        format: "json",
        ...requestParams,
      }),
  };
  resourcePools = {
    /**
     * @description Create ResourcePools
     *
     * @tags ResourcePools
     * @name CreateResourcePools
     * @summary Create ResourcePools
     * @request POST:/resource-pools/create
     * @secure
     */
    createResourcePools: (params: ModelsResourcePools, requestParams: RequestParams = {}) =>
      this.request<ModelsResourcePools, GinH>({
        path: `/resource-pools/create`,
        method: "POST",
        body: params,
        secure: true,
        type: ContentType.Json,
        format: "json",
        ...requestParams,
      }),

    /**
     * @description Get ResourcePools by offset
     *
     * @tags ResourcePools
     * @name GetResourcePoolsByOffset
     * @summary Get ResourcePools by offset
     * @request GET:/resource-pools/get-by-offset
     * @secure
     */
    getResourcePoolsByOffset: (query: { offset: number; limit: number }, params: RequestParams = {}) =>
      this.request<TypesArrayResponse, GinH>({
        path: `/resource-pools/get-by-offset`,
        method: "GET",
        query: query,
        secure: true,
        format: "json",
        ...params,
      }),

    /**
     * @description Get ResourcePools by uuid
     *
     * @tags ResourcePools
     * @name GetResourcePoolsByUuid
     * @summary Get ResourcePools by uuid
     * @request GET:/resource-pools/get/{uuid}
     * @secure
     */
    getResourcePoolsByUuid: (uuid: string, params: RequestParams = {}) =>
      this.request<ModelsResourcePools, GinH>({
        path: `/resource-pools/get/${uuid}`,
        method: "GET",
        secure: true,
        format: "json",
        ...params,
      }),

    /**
     * @description Hard Delete ResourcePools
     *
     * @tags ResourcePools
     * @name HardDeleteResourcePools
     * @summary Hard Delete ResourcePools
     * @request DELETE:/resource-pools/hard-delete/{uuid}
     * @secure
     */
    hardDeleteResourcePools: (uuid: string, params: RequestParams = {}) =>
      this.request<string, GinH>({
        path: `/resource-pools/hard-delete/${uuid}`,
        method: "DELETE",
        secure: true,
        format: "json",
        ...params,
      }),

    /**
     * @description Soft Delete ResourcePools
     *
     * @tags ResourcePools
     * @name SoftDeleteResourcePools
     * @summary Soft Delete ResourcePools
     * @request DELETE:/resource-pools/soft-delete/{uuid}
     * @secure
     */
    softDeleteResourcePools: (uuid: string, params: RequestParams = {}) =>
      this.request<string, GinH>({
        path: `/resource-pools/soft-delete/${uuid}`,
        method: "DELETE",
        secure: true,
        format: "json",
        ...params,
      }),

    /**
     * @description Update ResourcePools
     *
     * @tags ResourcePools
     * @name UpdateResourcePools
     * @summary Update ResourcePools
     * @request POST:/resource-pools/update
     * @secure
     */
    updateResourcePools: (params: ModelsResourcePools, requestParams: RequestParams = {}) =>
      this.request<ModelsResourcePools, GinH>({
        path: `/resource-pools/update`,
        method: "POST",
        body: params,
        secure: true,
        type: ContentType.Json,
        format: "json",
        ...requestParams,
      }),
  };
  resources = {
    /**
     * @description Create Resources
     *
     * @tags Resources
     * @name CreateResources
     * @summary Create Resources
     * @request POST:/resources/create
     * @secure
     */
    createResources: (params: ModelsResources, requestParams: RequestParams = {}) =>
      this.request<ModelsResources, GinH>({
        path: `/resources/create`,
        method: "POST",
        body: params,
        secure: true,
        type: ContentType.Json,
        format: "json",
        ...requestParams,
      }),

    /**
     * @description Get Resources by offset
     *
     * @tags Resources
     * @name GetResourcesByOffset
     * @summary Get Resources by offset
     * @request GET:/resources/get-by-offset
     * @secure
     */
    getResourcesByOffset: (query: { offset: number; limit: number }, params: RequestParams = {}) =>
      this.request<TypesArrayResponse, GinH>({
        path: `/resources/get-by-offset`,
        method: "GET",
        query: query,
        secure: true,
        format: "json",
        ...params,
      }),

    /**
     * @description Get Resources by uuid
     *
     * @tags Resources
     * @name GetResourcesByUuid
     * @summary Get Resources by uuid
     * @request GET:/resources/get/{uuid}
     * @secure
     */
    getResourcesByUuid: (uuid: string, params: RequestParams = {}) =>
      this.request<ModelsResources, GinH>({
        path: `/resources/get/${uuid}`,
        method: "GET",
        secure: true,
        format: "json",
        ...params,
      }),

    /**
     * @description Hard Delete Resources
     *
     * @tags Resources
     * @name HardDeleteResources
     * @summary Hard Delete Resources
     * @request DELETE:/resources/hard-delete/{uuid}
     * @secure
     */
    hardDeleteResources: (uuid: string, params: RequestParams = {}) =>
      this.request<string, GinH>({
        path: `/resources/hard-delete/${uuid}`,
        method: "DELETE",
        secure: true,
        format: "json",
        ...params,
      }),

    /**
     * @description Soft Delete Resources
     *
     * @tags Resources
     * @name SoftDeleteResources
     * @summary Soft Delete Resources
     * @request DELETE:/resources/soft-delete/{uuid}
     * @secure
     */
    softDeleteResources: (uuid: string, params: RequestParams = {}) =>
      this.request<string, GinH>({
        path: `/resources/soft-delete/${uuid}`,
        method: "DELETE",
        secure: true,
        format: "json",
        ...params,
      }),

    /**
     * @description Update Resources
     *
     * @tags Resources
     * @name UpdateResources
     * @summary Update Resources
     * @request POST:/resources/update
     * @secure
     */
    updateResources: (params: ModelsResources, requestParams: RequestParams = {}) =>
      this.request<ModelsResources, GinH>({
        path: `/resources/update`,
        method: "POST",
        body: params,
        secure: true,
        type: ContentType.Json,
        format: "json",
        ...requestParams,
      }),
  };
  roles = {
    /**
     * @description Create Roles
     *
     * @tags Roles
     * @name CreateRoles
     * @summary Create Roles
     * @request POST:/roles/create
     * @secure
     */
    createRoles: (params: ModelsRoles, requestParams: RequestParams = {}) =>
      this.request<ModelsRoles, GinH>({
        path: `/roles/create`,
        method: "POST",
        body: params,
        secure: true,
        type: ContentType.Json,
        format: "json",
        ...requestParams,
      }),

    /**
     * @description Get Roles by offset
     *
     * @tags Roles
     * @name GetRolesByOffset
     * @summary Get Roles by offset
     * @request GET:/roles/get-by-offset
     * @secure
     */
    getRolesByOffset: (query: { offset: number; limit: number }, params: RequestParams = {}) =>
      this.request<TypesArrayResponse, GinH>({
        path: `/roles/get-by-offset`,
        method: "GET",
        query: query,
        secure: true,
        format: "json",
        ...params,
      }),

    /**
     * @description Get Roles by uuid
     *
     * @tags Roles
     * @name GetRolesByUuid
     * @summary Get Roles by uuid
     * @request GET:/roles/get/{uuid}
     * @secure
     */
    getRolesByUuid: (uuid: string, params: RequestParams = {}) =>
      this.request<ModelsRoles, GinH>({
        path: `/roles/get/${uuid}`,
        method: "GET",
        secure: true,
        format: "json",
        ...params,
      }),

    /**
     * @description Hard Delete Roles
     *
     * @tags Roles
     * @name HardDeleteRoles
     * @summary Hard Delete Roles
     * @request DELETE:/roles/hard-delete/{uuid}
     * @secure
     */
    hardDeleteRoles: (uuid: string, params: RequestParams = {}) =>
      this.request<string, GinH>({
        path: `/roles/hard-delete/${uuid}`,
        method: "DELETE",
        secure: true,
        format: "json",
        ...params,
      }),

    /**
     * @description Soft Delete Roles
     *
     * @tags Roles
     * @name SoftDeleteRoles
     * @summary Soft Delete Roles
     * @request DELETE:/roles/soft-delete/{uuid}
     * @secure
     */
    softDeleteRoles: (uuid: string, params: RequestParams = {}) =>
      this.request<string, GinH>({
        path: `/roles/soft-delete/${uuid}`,
        method: "DELETE",
        secure: true,
        format: "json",
        ...params,
      }),

    /**
     * @description Update Roles
     *
     * @tags Roles
     * @name UpdateRoles
     * @summary Update Roles
     * @request POST:/roles/update
     * @secure
     */
    updateRoles: (params: ModelsRoles, requestParams: RequestParams = {}) =>
      this.request<ModelsRoles, GinH>({
        path: `/roles/update`,
        method: "POST",
        body: params,
        secure: true,
        type: ContentType.Json,
        format: "json",
        ...requestParams,
      }),
  };
  serviceVersions = {
    /**
     * @description Create ServiceVersions
     *
     * @tags ServiceVersions
     * @name CreateServiceVersions
     * @summary Create ServiceVersions
     * @request POST:/service-versions/create
     * @secure
     */
    createServiceVersions: (params: ModelsServiceVersions, requestParams: RequestParams = {}) =>
      this.request<ModelsServiceVersions, GinH>({
        path: `/service-versions/create`,
        method: "POST",
        body: params,
        secure: true,
        type: ContentType.Json,
        format: "json",
        ...requestParams,
      }),

    /**
     * @description Get ServiceVersions by offset
     *
     * @tags ServiceVersions
     * @name GetServiceVersionsByOffset
     * @summary Get ServiceVersions by offset
     * @request GET:/service-versions/get-by-offset
     * @secure
     */
    getServiceVersionsByOffset: (query: { offset: number; limit: number }, params: RequestParams = {}) =>
      this.request<TypesArrayResponse, GinH>({
        path: `/service-versions/get-by-offset`,
        method: "GET",
        query: query,
        secure: true,
        format: "json",
        ...params,
      }),

    /**
     * @description Get ServiceVersions by sid
     *
     * @tags ServiceVersions
     * @name GetServiceVersionsBySid
     * @summary Get ServiceVersions by sid
     * @request GET:/service-versions/get-by-sid
     * @secure
     */
    getServiceVersionsBySid: (query: { sid: number; offset: number; limit: number }, params: RequestParams = {}) =>
      this.request<TypesArrayResponse, GinH>({
        path: `/service-versions/get-by-sid`,
        method: "GET",
        query: query,
        secure: true,
        format: "json",
        ...params,
      }),

    /**
     * @description Get ServiceVersions by uuid
     *
     * @tags ServiceVersions
     * @name GetServiceVersionsByUuid
     * @summary Get ServiceVersions by uuid
     * @request GET:/service-versions/get/{uuid}
     * @secure
     */
    getServiceVersionsByUuid: (uuid: string, params: RequestParams = {}) =>
      this.request<ModelsServiceVersions, GinH>({
        path: `/service-versions/get/${uuid}`,
        method: "GET",
        secure: true,
        format: "json",
        ...params,
      }),

    /**
     * @description Hard Delete ServiceVersions
     *
     * @tags ServiceVersions
     * @name HardDeleteServiceVersions
     * @summary Hard Delete ServiceVersions
     * @request DELETE:/service-versions/hard-delete/{uuid}
     * @secure
     */
    hardDeleteServiceVersions: (uuid: string, params: RequestParams = {}) =>
      this.request<string, GinH>({
        path: `/service-versions/hard-delete/${uuid}`,
        method: "DELETE",
        secure: true,
        format: "json",
        ...params,
      }),

    /**
     * @description Soft Delete ServiceVersions
     *
     * @tags ServiceVersions
     * @name SoftDeleteServiceVersions
     * @summary Soft Delete ServiceVersions
     * @request DELETE:/service-versions/soft-delete/{uuid}
     * @secure
     */
    softDeleteServiceVersions: (uuid: string, params: RequestParams = {}) =>
      this.request<string, GinH>({
        path: `/service-versions/soft-delete/${uuid}`,
        method: "DELETE",
        secure: true,
        format: "json",
        ...params,
      }),

    /**
     * @description Update ServiceVersions
     *
     * @tags ServiceVersions
     * @name UpdateServiceVersions
     * @summary Update ServiceVersions
     * @request POST:/service-versions/update
     * @secure
     */
    updateServiceVersions: (params: ModelsServiceVersions, requestParams: RequestParams = {}) =>
      this.request<ModelsServiceVersions, GinH>({
        path: `/service-versions/update`,
        method: "POST",
        body: params,
        secure: true,
        type: ContentType.Json,
        format: "json",
        ...requestParams,
      }),
  };
  services = {
    /**
     * @description Create Services
     *
     * @tags Services
     * @name CreateServices
     * @summary Create Services
     * @request POST:/services/create
     * @secure
     */
    createServices: (params: ModelsServices, requestParams: RequestParams = {}) =>
      this.request<ModelsServices, GinH>({
        path: `/services/create`,
        method: "POST",
        body: params,
        secure: true,
        type: ContentType.Json,
        format: "json",
        ...requestParams,
      }),

    /**
     * @description Get Services by offset
     *
     * @tags Services
     * @name GetServicesByOffset
     * @summary Get Services by offset
     * @request GET:/services/get-by-offset
     * @secure
     */
    getServicesByOffset: (query: { offset: number; limit: number }, params: RequestParams = {}) =>
      this.request<TypesArrayResponse, GinH>({
        path: `/services/get-by-offset`,
        method: "GET",
        query: query,
        secure: true,
        format: "json",
        ...params,
      }),

    /**
     * @description Get Services by uuid
     *
     * @tags Services
     * @name GetServicesByUuid
     * @summary Get Services by uuid
     * @request GET:/services/get/{uuid}
     * @secure
     */
    getServicesByUuid: (uuid: string, params: RequestParams = {}) =>
      this.request<ModelsServices, GinH>({
        path: `/services/get/${uuid}`,
        method: "GET",
        secure: true,
        format: "json",
        ...params,
      }),

    /**
     * @description Hard Delete Services
     *
     * @tags Services
     * @name HardDeleteServices
     * @summary Hard Delete Services
     * @request DELETE:/services/hard-delete/{uuid}
     * @secure
     */
    hardDeleteServices: (uuid: string, params: RequestParams = {}) =>
      this.request<string, GinH>({
        path: `/services/hard-delete/${uuid}`,
        method: "DELETE",
        secure: true,
        format: "json",
        ...params,
      }),

    /**
     * @description Soft Delete Services
     *
     * @tags Services
     * @name SoftDeleteServices
     * @summary Soft Delete Services
     * @request DELETE:/services/soft-delete/{uuid}
     * @secure
     */
    softDeleteServices: (uuid: string, params: RequestParams = {}) =>
      this.request<string, GinH>({
        path: `/services/soft-delete/${uuid}`,
        method: "DELETE",
        secure: true,
        format: "json",
        ...params,
      }),

    /**
     * @description Update Services
     *
     * @tags Services
     * @name UpdateServices
     * @summary Update Services
     * @request POST:/services/update
     * @secure
     */
    updateServices: (params: ModelsServices, requestParams: RequestParams = {}) =>
      this.request<ModelsServices, GinH>({
        path: `/services/update`,
        method: "POST",
        body: params,
        secure: true,
        type: ContentType.Json,
        format: "json",
        ...requestParams,
      }),
  };
  setup = {
    /**
     * @description Setup system
     *
     * @tags Setup
     * @name SetupSystem
     * @summary Setup system
     * @request POST:/setup
     * @secure
     */
    setupSystem: (params: TypesSetupParams, requestParams: RequestParams = {}) =>
      this.request<ModelsUsers, GinH>({
        path: `/setup`,
        method: "POST",
        body: params,
        secure: true,
        type: ContentType.Json,
        format: "json",
        ...requestParams,
      }),
  };
  subscriptions = {
    /**
     * @description Create Subscriptions
     *
     * @tags Subscriptions
     * @name CreateSubscriptions
     * @summary Create Subscriptions
     * @request POST:/subscriptions/create
     * @secure
     */
    createSubscriptions: (params: ModelsSubscriptions, requestParams: RequestParams = {}) =>
      this.request<ModelsSubscriptions, GinH>({
        path: `/subscriptions/create`,
        method: "POST",
        body: params,
        secure: true,
        type: ContentType.Json,
        format: "json",
        ...requestParams,
      }),

    /**
     * @description Get Subscriptions by offset
     *
     * @tags Subscriptions
     * @name GetSubscriptionsByOffset
     * @summary Get Subscriptions by offset
     * @request GET:/subscriptions/get-by-offset
     * @secure
     */
    getSubscriptionsByOffset: (query: { offset: number; limit: number }, params: RequestParams = {}) =>
      this.request<TypesArrayResponse, GinH>({
        path: `/subscriptions/get-by-offset`,
        method: "GET",
        query: query,
        secure: true,
        format: "json",
        ...params,
      }),

    /**
     * @description Get Subscriptions by uuid
     *
     * @tags Subscriptions
     * @name GetSubscriptionsByUuid
     * @summary Get Subscriptions by uuid
     * @request GET:/subscriptions/get/{uuid}
     * @secure
     */
    getSubscriptionsByUuid: (uuid: string, params: RequestParams = {}) =>
      this.request<ModelsSubscriptions, GinH>({
        path: `/subscriptions/get/${uuid}`,
        method: "GET",
        secure: true,
        format: "json",
        ...params,
      }),

    /**
     * @description Hard Delete Subscriptions
     *
     * @tags Subscriptions
     * @name HardDeleteSubscriptions
     * @summary Hard Delete Subscriptions
     * @request DELETE:/subscriptions/hard-delete/{uuid}
     * @secure
     */
    hardDeleteSubscriptions: (uuid: string, params: RequestParams = {}) =>
      this.request<string, GinH>({
        path: `/subscriptions/hard-delete/${uuid}`,
        method: "DELETE",
        secure: true,
        format: "json",
        ...params,
      }),

    /**
     * @description Soft Delete Subscriptions
     *
     * @tags Subscriptions
     * @name SoftDeleteSubscriptions
     * @summary Soft Delete Subscriptions
     * @request DELETE:/subscriptions/soft-delete/{uuid}
     * @secure
     */
    softDeleteSubscriptions: (uuid: string, params: RequestParams = {}) =>
      this.request<string, GinH>({
        path: `/subscriptions/soft-delete/${uuid}`,
        method: "DELETE",
        secure: true,
        format: "json",
        ...params,
      }),

    /**
     * @description Update Subscriptions
     *
     * @tags Subscriptions
     * @name UpdateSubscriptions
     * @summary Update Subscriptions
     * @request POST:/subscriptions/update
     * @secure
     */
    updateSubscriptions: (params: ModelsSubscriptions, requestParams: RequestParams = {}) =>
      this.request<ModelsSubscriptions, GinH>({
        path: `/subscriptions/update`,
        method: "POST",
        body: params,
        secure: true,
        type: ContentType.Json,
        format: "json",
        ...requestParams,
      }),
  };
  users = {
    /**
     * @description BYO password Users
     *
     * @tags Users
     * @name ByoPasswdUsers
     * @summary BYO password Users
     * @request POST:/users/byo-password
     * @secure
     */
    byoPasswdUsers: (params: TypesUserChangePasswdParams, requestParams: RequestParams = {}) =>
      this.request<string, GinH>({
        path: `/users/byo-password`,
        method: "POST",
        body: params,
        secure: true,
        type: ContentType.Json,
        format: "json",
        ...requestParams,
      }),

    /**
     * @description Change password Users
     *
     * @tags Users
     * @name ChangePasswdUsers
     * @summary Change password Users
     * @request POST:/users/change-password
     * @secure
     */
    changePasswdUsers: (params: TypesUserChangePasswdParams, requestParams: RequestParams = {}) =>
      this.request<string, GinH>({
        path: `/users/change-password`,
        method: "POST",
        body: params,
        secure: true,
        type: ContentType.Json,
        format: "json",
        ...requestParams,
      }),

    /**
     * @description Create Users
     *
     * @tags Users
     * @name CreateUsers
     * @summary Create Users
     * @request POST:/users/create
     * @secure
     */
    createUsers: (params: ModelsUsers, requestParams: RequestParams = {}) =>
      this.request<ModelsUsers, GinH>({
        path: `/users/create`,
        method: "POST",
        body: params,
        secure: true,
        type: ContentType.Json,
        format: "json",
        ...requestParams,
      }),

    /**
     * @description Get Users by offset
     *
     * @tags Users
     * @name GetUsersByOffset
     * @summary Get Users by offset
     * @request GET:/users/get-by-offset
     * @secure
     */
    getUsersByOffset: (query: { offset: number; limit: number }, params: RequestParams = {}) =>
      this.request<TypesArrayResponse, GinH>({
        path: `/users/get-by-offset`,
        method: "GET",
        query: query,
        secure: true,
        format: "json",
        ...params,
      }),

    /**
     * @description Get Users by uuid
     *
     * @tags Users
     * @name GetUsersByUuid
     * @summary Get Users by uuid
     * @request GET:/users/get/{uuid}
     * @secure
     */
    getUsersByUuid: (uuid: string, params: RequestParams = {}) =>
      this.request<ModelsUsers, GinH>({
        path: `/users/get/${uuid}`,
        method: "GET",
        secure: true,
        format: "json",
        ...params,
      }),

    /**
     * @description Hard Delete Users
     *
     * @tags Users
     * @name HardDeleteUsers
     * @summary Hard Delete Users
     * @request DELETE:/users/hard-delete/{uuid}
     * @secure
     */
    hardDeleteUsers: (uuid: string, params: RequestParams = {}) =>
      this.request<string, GinH>({
        path: `/users/hard-delete/${uuid}`,
        method: "DELETE",
        secure: true,
        format: "json",
        ...params,
      }),

    /**
     * @description Reset password Users
     *
     * @tags Users
     * @name ResetPasswdUsers
     * @summary Reset password Users
     * @request POST:/users/reset-password
     * @secure
     */
    resetPasswdUsers: (params: TypesUserResetPasswdParams, requestParams: RequestParams = {}) =>
      this.request<ModelsUsers, GinH>({
        path: `/users/reset-password`,
        method: "POST",
        body: params,
        secure: true,
        type: ContentType.Json,
        format: "json",
        ...requestParams,
      }),

    /**
     * @description Soft Delete Users
     *
     * @tags Users
     * @name SoftDeleteUsers
     * @summary Soft Delete Users
     * @request DELETE:/users/soft-delete/{uuid}
     * @secure
     */
    softDeleteUsers: (uuid: string, params: RequestParams = {}) =>
      this.request<string, GinH>({
        path: `/users/soft-delete/${uuid}`,
        method: "DELETE",
        secure: true,
        format: "json",
        ...params,
      }),

    /**
     * @description Update Users
     *
     * @tags Users
     * @name UpdateUsers
     * @summary Update Users
     * @request POST:/users/update
     * @secure
     */
    updateUsers: (params: ModelsUsers, requestParams: RequestParams = {}) =>
      this.request<ModelsUsers, GinH>({
        path: `/users/update`,
        method: "POST",
        body: params,
        secure: true,
        type: ContentType.Json,
        format: "json",
        ...requestParams,
      }),
  };
}
