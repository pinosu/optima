use antex::ColorMode;
use dsntk_common::Jsonify;
use dsntk_evaluator::evaluate_context;
use dsntk_feel::FeelScope;
use dsntk_workspace::Workspaces;
use once_cell::sync::Lazy;
use std::ffi::CStr;
use std::os::raw::c_char;
use std::sync::RwLock;
use std::{ptr, slice};

struct App {
  workspaces: Option<Workspaces>,
}

impl App {
  fn new() -> Self {
    App { workspaces: None }
  }
}

static APP: Lazy<RwLock<App>> = Lazy::new(|| RwLock::new(App::new()));

/// Loads decision models.
#[no_mangle]
pub unsafe extern "C" fn load_models(c_dir: *const c_char) -> usize {
  let dir_str = unsafe { CStr::from_ptr(c_dir) };
  if let Ok(dir) = dir_str.to_str() {
    if let Ok(mut app) = APP.write() {
      let workspaces = Workspaces::new(&[dir.into()], ColorMode::Off, true);
      app.workspaces = Some(workspaces);
      return 1;
    }
  }
  0 // error code
}

/// Evaluates the invocable.
#[no_mangle]
pub unsafe extern "C" fn evaluate_invocable(c_invocable_name: *const c_char, c_input_data: *const c_char, output_data_len: *mut usize) -> *mut u8 {
  let invocable_name_str = unsafe { CStr::from_ptr(c_invocable_name) };
  let input_data_str = unsafe { CStr::from_ptr(c_input_data) };
  if let Ok(invocable_name) = invocable_name_str.to_str() {
    if let Ok(input_data) = input_data_str.to_str() {
      if let Ok(workspaces) = APP.read() {
        if let Some(workspaces) = &workspaces.workspaces {
          if let Ok(ctx) = evaluate_context(&FeelScope::default(), input_data) {
            if let Ok(result) = workspaces.evaluate(invocable_name, &ctx) {
              let json_string = result.jsonify().as_bytes().to_vec();
              *output_data_len = json_string.len();
              return Box::into_raw(json_string.into_boxed_slice()) as *mut u8;
            }
          }
        }
      }
    }
  }
  *output_data_len = 0; // error code
  ptr::null_mut()
}

/// Free the memory allocated in Rust, but the raw pointer was returned to Go.
#[no_mangle]
pub unsafe extern "C" fn free_memory(ptr: *mut u8, len: usize) {
  if !ptr.is_null() {
    drop(Box::from_raw(slice::from_raw_parts_mut(ptr, len)));
  }
}