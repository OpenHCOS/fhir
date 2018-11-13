extern crate wasm_bindgen;
use wasm_bindgen::prelude::*;
mod proto;
use self::proto::addressbook::{Person};

extern crate wbg_rand;
use wbg_rand::{Rng, wasm_rng, math_random_rng};

#[wasm_bindgen]
pub fn greet(name: &str) -> String {
    // rand panic when use wasm-bindgen · Issue #616 · rust-random/rand  https://github.com/rust-random/rand/issues/616
    let v = Person::new();
    let x: usize = wasm_rng().gen_range(0, 100);
    format!("Hello OpenHCOS , {} , {} !", name, x)
}