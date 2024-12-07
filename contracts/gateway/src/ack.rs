use cosmwasm_schema::cw_serde;
use cosmwasm_std::{to_json_binary, Binary};

#[cw_serde]
pub enum Ack {
    Result(Binary),
    Error(String),
}

pub fn make_ack_success() -> Binary {
    let res = Ack::Result(b"1".into());
    to_json_binary(&res).unwrap()
}

pub fn make_ack_fail(err: String) -> Binary {
    let res = Ack::Error(err);
    to_json_binary(&res).unwrap()
}