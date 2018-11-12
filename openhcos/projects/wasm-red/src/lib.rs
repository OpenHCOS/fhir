extern crate wasm_bindgen;
use wasm_bindgen::prelude::*;
// extern crate rand;
// use rand::prelude::*;
// use rand::rngs::OsRng;
// use rand::distributions::StandardNormal;

extern crate wbg_rand;
use wbg_rand::{Rng, wasm_rng, math_random_rng};

#[wasm_bindgen]
pub fn greet(name: &str) -> String {
    // rand panic when use wasm-bindgen · Issue #616 · rust-random/rand  https://github.com/rust-random/rand/issues/616
    //let mut small_rng = SmallRng::from_entropy();
    //let val: f64 = small_rng.sample(StandardNormal);
    // basic usage with random():
    // let x: u8 = random();

    let x: usize = wasm_rng().gen_range(0, 1000);
    format!("Hello OpenHCOS , {} , {} !", name, x)
}