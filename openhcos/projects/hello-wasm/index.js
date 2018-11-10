const w = import("./hello_wasm");

w.then(js => {
  js.greet("World!");
});