const w = import("./wasm_red");

w.then(js => {
  console.log(js.greet("World!"));
});