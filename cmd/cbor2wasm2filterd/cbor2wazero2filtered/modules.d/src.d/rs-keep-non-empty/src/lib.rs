use std::sync::RwLock;

use ciborium::Value;

static INPUT: RwLock<Vec<u8>> = RwLock::new(vec![]);

pub fn _initialize_input_buffer(sz: u64, init: u64) -> Result<*mut u8, &'static str> {
    let mut guard = INPUT.try_write().map_err(|_| "unable to write lock")?;
    let mv: &mut Vec<u8> = &mut guard;
    mv.resize(sz as usize, init as u8);
    Ok(mv.as_mut_ptr())
}

#[allow(unsafe_code)]
#[no_mangle]
pub extern "C" fn initialize_input_buffer(sz: u64, init: u64) -> *mut u8 {
    _initialize_input_buffer(sz, init)
        .ok()
        .unwrap_or_else(std::ptr::null_mut)
}

pub fn filter_slice(s: &[u8]) -> Result<bool, &'static str> {
    let val: Value = ciborium::from_reader(s).map_err(|_| "invalid value")?;
    match val {
        Value::Array(v) => {
            let empty: bool = v.is_empty();
            let keep: bool = !empty;
            Ok(keep)
        }
        _ => Err("invalid array"),
    }
}

pub fn _filter() -> Result<u64, &'static str> {
    let guard = INPUT.try_read().map_err(|_| "unable to read lock")?;
    filter_slice(&guard).map(|b| match b {
        true => -1_i64 as u64,
        false => 0,
    })
}

#[allow(unsafe_code)]
#[no_mangle]
pub fn filter() -> u64 {
    _filter().ok().unwrap_or(0)
}
