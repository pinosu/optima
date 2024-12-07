use cosmwasm_schema::{cw_serde, QueryResponses};
use cosmwasm_std::CustomMsg;

#[cw_serde]
pub struct InstantiateMsg {}

#[cw_serde]
pub enum ExecuteMsg {
    Evaluate {
        invocable_name: String,
        input_data: String,
    }
}

#[cw_serde]
#[derive(QueryResponses)]
pub enum QueryMsg {}

#[cw_serde]
pub enum OptimaMsg {
    Invocable {
        invocable_name: String,
        input_data: String,
    }
}

impl CustomMsg for OptimaMsg {}
