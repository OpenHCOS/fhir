extern crate protoc_rust;

use protoc_rust::Customize;

// https://github.com/hyperledger/sawtooth-sabre/blob/master/example/intkey_multiply/processor/build.rs
// https://github.com/hyperledger/sawtooth-sabre/blob/master/tp/build.rs
use std::fs;
use std::fs::File;
use std::io::prelude::*;

fn main() {
	protoc_rust::run(protoc_rust::Args {
	    out_dir: "src/proto",
	    input: &["proto/addressbook.proto",
		"../../../proto/stu3/datatypes.proto",
		"../../../proto/stu3/codes.proto"
		],
	    includes: &["proto","C:/go/src","../../../"],
	    customize: Customize {
	      ..Default::default()
	    },
	}).expect("protoc");

    let mut file = File::create("src/proto/mod.rs").unwrap();
    file.write_all(b"pub mod addressbook;\n").unwrap();
    file.write_all(b"pub mod codes;\n").unwrap();
    file.write_all(b"pub mod datatypes;\n").unwrap();
}