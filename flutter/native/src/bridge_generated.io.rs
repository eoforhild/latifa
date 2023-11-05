use super::*;
// Section: wire functions

#[no_mangle]
pub extern "C" fn wire_kdf(port_: i64, km: *mut wire_uint_8_list) {
    wire_kdf_impl(port_, km)
}

#[no_mangle]
pub extern "C" fn wire_generate_keys_and_dump(port_: i64) {
    wire_generate_keys_and_dump_impl(port_)
}

#[no_mangle]
pub extern "C" fn wire_register_and_publish(port_: i64, reg_form: *mut wire_uint_8_list) {
    wire_register_and_publish_impl(port_, reg_form)
}

#[no_mangle]
pub extern "C" fn wire_login(port_: i64, log_form: *mut wire_uint_8_list) {
    wire_login_impl(port_, log_form)
}

#[no_mangle]
pub extern "C" fn wire_logout(port_: i64) {
    wire_logout_impl(port_)
}

#[no_mangle]
pub extern "C" fn wire_request_connection(port_: i64, email: *mut wire_uint_8_list) {
    wire_request_connection_impl(port_, email)
}

#[no_mangle]
pub extern "C" fn wire_pending_requests(port_: i64) {
    wire_pending_requests_impl(port_)
}

#[no_mangle]
pub extern "C" fn wire_approved_requests(port_: i64) {
    wire_approved_requests_impl(port_)
}

#[no_mangle]
pub extern "C" fn wire_fetch_keys_handshake(port_: i64, req_id: *mut wire_uint_8_list) {
    wire_fetch_keys_handshake_impl(port_, req_id)
}

// Section: allocate functions

#[no_mangle]
pub extern "C" fn new_uint_8_list_0(len: i32) -> *mut wire_uint_8_list {
    let ans = wire_uint_8_list {
        ptr: support::new_leak_vec_ptr(Default::default(), len),
        len,
    };
    support::new_leak_box_ptr(ans)
}

// Section: related functions

// Section: impl Wire2Api

impl Wire2Api<String> for *mut wire_uint_8_list {
    fn wire2api(self) -> String {
        let vec: Vec<u8> = self.wire2api();
        String::from_utf8_lossy(&vec).into_owned()
    }
}

impl Wire2Api<Vec<u8>> for *mut wire_uint_8_list {
    fn wire2api(self) -> Vec<u8> {
        unsafe {
            let wrap = support::box_from_leak_ptr(self);
            support::vec_from_leak_ptr(wrap.ptr, wrap.len)
        }
    }
}
// Section: wire structs

#[repr(C)]
#[derive(Clone)]
pub struct wire_uint_8_list {
    ptr: *mut u8,
    len: i32,
}

// Section: impl NewWithNullPtr

pub trait NewWithNullPtr {
    fn new_with_null_ptr() -> Self;
}

impl<T> NewWithNullPtr for *mut T {
    fn new_with_null_ptr() -> Self {
        std::ptr::null_mut()
    }
}

// Section: sync execution mode utility

#[no_mangle]
pub extern "C" fn free_WireSyncReturn(ptr: support::WireSyncReturn) {
    unsafe {
        let _ = support::box_from_leak_ptr(ptr);
    };
}
