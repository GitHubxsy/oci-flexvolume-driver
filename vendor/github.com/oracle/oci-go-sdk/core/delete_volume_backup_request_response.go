// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package core

import (
    "github.com/oracle/oci-go-sdk/common"
    "net/http"
)

// DeleteVolumeBackupRequest wrapper for the DeleteVolumeBackup operation
type DeleteVolumeBackupRequest struct {
        
 // The OCID of the volume backup. 
        VolumeBackupId *string `mandatory:"true" contributesTo:"path" name:"volumeBackupId"`
        
 // For optimistic concurrency control. In the PUT or DELETE call for a resource, set the `if-match`
 // parameter to the value of the etag from a previous GET or POST response for that resource.  The resource
 // will be updated or deleted only if the etag you provide matches the resource's current etag value. 
        IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`
}

func (request DeleteVolumeBackupRequest) String() string {
    return common.PointerString(request)
}

// DeleteVolumeBackupResponse wrapper for the DeleteVolumeBackup operation
type DeleteVolumeBackupResponse struct {

    // The underlying http response
    RawResponse *http.Response

    
 // Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
 // a particular request, please provide the request ID.
    OpcRequestId *string `presentIn:"header" name:"opc-request-id"`


}

func (response DeleteVolumeBackupResponse) String() string {
    return common.PointerString(response)
}

