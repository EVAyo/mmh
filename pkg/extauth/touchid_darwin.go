// +build darwin

package extauth

/*
#cgo CFLAGS: -x objective-c -fmodules -fblocks
#cgo LDFLAGS: -framework CoreFoundation -framework LocalAuthentication -framework Foundation
#include <stdlib.h>
#include <stdio.h>
#import <LocalAuthentication/LocalAuthentication.h>
int Authenticate(char const* reason) {
 LAContext *laContext = [[LAContext alloc] init];
 NSError *authError = nil;
 dispatch_semaphore_t sema = dispatch_semaphore_create(0);
 NSString *nsReason = [NSString stringWithUTF8String:reason];
 __block int result = 0;
 if ([laContext canEvaluatePolicy:LAPolicyDeviceOwnerAuthentication error:&authError]) {
   [laContext evaluatePolicy:LAPolicyDeviceOwnerAuthentication
	 localizedReason:nsReason
	 reply:^(BOOL success, NSError *error) {
	   if (success) {
		 result = 1;
	   } else {
		 result = 2;
	   }
	   dispatch_semaphore_signal(sema);
	 }];
 }
 dispatch_semaphore_wait(sema, DISPATCH_TIME_FOREVER);
 dispatch_release(sema);
 return result;
}
*/
import (
	"C"
)
import (
	"errors"
	"unsafe"
)

func TouchIDAuth(reason string) (bool, error) {
	reasonStr := C.CString(reason)
	defer C.free(unsafe.Pointer(reasonStr))

	result := C.Authenticate(reasonStr)
	switch result {
	case 1:
		return true, nil
	case 2:
		return false, nil
	}

	return false, errors.New("error occurred accessing biometrics")
}
