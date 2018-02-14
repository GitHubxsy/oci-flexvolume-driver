// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package ffsw

import (
    "github.com/oracle/oci-go-sdk/common"
    "net/http"
)

// ListSnapshotsRequest wrapper for the ListSnapshots operation
type ListSnapshotsRequest struct {
        
 // The OCID of the file system. This feature is currently in preview and may change before public release. Do not use it for production workloads.
        FileSystemId *string `mandatory:"true" contributesTo:"query" name:"fileSystemId"`
        
 // The maximum number of items to return in a paginated "List" call.
 // Example: `500` 
        Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`
        
 // The value of the `opc-next-page` response header from the previous "List" call. 
        Page *string `mandatory:"false" contributesTo:"query" name:"page"`
        
 // Snapshot name.
 // Example: `sunday` 
        Name *string `mandatory:"false" contributesTo:"query" name:"name"`
        
 // Filter results by the specified lifecycle state. Must be a valid
 // state for the resource type. 
        LifecycleState ListSnapshotsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`
        
 // Filter results by OCID. Must be an OCID of the correct type for
 // the resouce type. 
        Id *string `mandatory:"false" contributesTo:"query" name:"id"`
        
 // The field to sort by.  Only one sort order may be provided.
 // Time created is default ordered as descending.  Display name,
 // path, and name is default ordered as ascending. 
        SortBy ListSnapshotsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`
        
 // The sort order to use, either 'asc' or 'desc'. 
        SortOrder ListSnapshotsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`
}

func (request ListSnapshotsRequest) String() string {
    return common.PointerString(request)
}

// ListSnapshotsResponse wrapper for the ListSnapshots operation
type ListSnapshotsResponse struct {

    // The underlying http response
    RawResponse *http.Response
    
 // The []SnapshotSummary instance
    Items []SnapshotSummary `presentIn:"body"`

    
 // For pagination of a list of items. When paging through
 // a list, if this header appears in the response, then a
 // partial list might have been returned. Include this
 // value as the `page` parameter for the subsequent GET
 // request to get the next batch of items.
    OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
    
 // Unique Oracle-assigned identifier for the request. If
 // you need to contact Oracle about a particular request,
 // please provide the request ID.
    OpcRequestId *string `presentIn:"header" name:"opc-request-id"`


}

func (response ListSnapshotsResponse) String() string {
    return common.PointerString(response)
}

// ListSnapshotsLifecycleStateEnum Enum with underlying type: string
type ListSnapshotsLifecycleStateEnum string

// Set of constants representing the allowable values for ListSnapshotsLifecycleState
const (
    ListSnapshotsLifecycleStateCreating ListSnapshotsLifecycleStateEnum = "CREATING"
    ListSnapshotsLifecycleStateActive ListSnapshotsLifecycleStateEnum = "ACTIVE"
    ListSnapshotsLifecycleStateDeleting ListSnapshotsLifecycleStateEnum = "DELETING"
    ListSnapshotsLifecycleStateDeleted ListSnapshotsLifecycleStateEnum = "DELETED"
    ListSnapshotsLifecycleStateFailed ListSnapshotsLifecycleStateEnum = "FAILED"
    ListSnapshotsLifecycleStateUnknown ListSnapshotsLifecycleStateEnum = "UNKNOWN"
)

var mappingListSnapshotsLifecycleState = map[string]ListSnapshotsLifecycleStateEnum { 
    "CREATING": ListSnapshotsLifecycleStateCreating,
    "ACTIVE": ListSnapshotsLifecycleStateActive,
    "DELETING": ListSnapshotsLifecycleStateDeleting,
    "DELETED": ListSnapshotsLifecycleStateDeleted,
    "FAILED": ListSnapshotsLifecycleStateFailed,
    "UNKNOWN": ListSnapshotsLifecycleStateUnknown,
}

// GetListSnapshotsLifecycleStateEnumValues Enumerates the set of values for ListSnapshotsLifecycleState
func GetListSnapshotsLifecycleStateEnumValues() []ListSnapshotsLifecycleStateEnum {
   values := make([]ListSnapshotsLifecycleStateEnum, 0)
   for _, v := range mappingListSnapshotsLifecycleState {
      if v != ListSnapshotsLifecycleStateUnknown {
         values = append(values, v)
      }
   }
   return values
}

// ListSnapshotsSortByEnum Enum with underlying type: string
type ListSnapshotsSortByEnum string

// Set of constants representing the allowable values for ListSnapshotsSortBy
const (
    ListSnapshotsSortByTimecreated ListSnapshotsSortByEnum = "TIMECREATED"
    ListSnapshotsSortByDisplayname ListSnapshotsSortByEnum = "DISPLAYNAME"
    ListSnapshotsSortByName ListSnapshotsSortByEnum = "NAME"
    ListSnapshotsSortByPath ListSnapshotsSortByEnum = "PATH"
    ListSnapshotsSortByUnknown ListSnapshotsSortByEnum = "UNKNOWN"
)

var mappingListSnapshotsSortBy = map[string]ListSnapshotsSortByEnum { 
    "TIMECREATED": ListSnapshotsSortByTimecreated,
    "DISPLAYNAME": ListSnapshotsSortByDisplayname,
    "NAME": ListSnapshotsSortByName,
    "PATH": ListSnapshotsSortByPath,
    "UNKNOWN": ListSnapshotsSortByUnknown,
}

// GetListSnapshotsSortByEnumValues Enumerates the set of values for ListSnapshotsSortBy
func GetListSnapshotsSortByEnumValues() []ListSnapshotsSortByEnum {
   values := make([]ListSnapshotsSortByEnum, 0)
   for _, v := range mappingListSnapshotsSortBy {
      if v != ListSnapshotsSortByUnknown {
         values = append(values, v)
      }
   }
   return values
}

// ListSnapshotsSortOrderEnum Enum with underlying type: string
type ListSnapshotsSortOrderEnum string

// Set of constants representing the allowable values for ListSnapshotsSortOrder
const (
    ListSnapshotsSortOrderAsc ListSnapshotsSortOrderEnum = "ASC"
    ListSnapshotsSortOrderDesc ListSnapshotsSortOrderEnum = "DESC"
    ListSnapshotsSortOrderUnknown ListSnapshotsSortOrderEnum = "UNKNOWN"
)

var mappingListSnapshotsSortOrder = map[string]ListSnapshotsSortOrderEnum { 
    "ASC": ListSnapshotsSortOrderAsc,
    "DESC": ListSnapshotsSortOrderDesc,
    "UNKNOWN": ListSnapshotsSortOrderUnknown,
}

// GetListSnapshotsSortOrderEnumValues Enumerates the set of values for ListSnapshotsSortOrder
func GetListSnapshotsSortOrderEnumValues() []ListSnapshotsSortOrderEnum {
   values := make([]ListSnapshotsSortOrderEnum, 0)
   for _, v := range mappingListSnapshotsSortOrder {
      if v != ListSnapshotsSortOrderUnknown {
         values = append(values, v)
      }
   }
   return values
}

