use cosmwasm_schema::{cw_serde, QueryResponses};

#[cw_serde]
pub struct InstantiateMsg {}

#[cw_serde]
pub enum ExecuteMsg {
    Evaluate {
        channel: String,
        job_id: u64,
        invocable_name: String,
        input_data: String,
    }
}

#[cw_serde]
pub enum IbcExecuteMsg {
    Evaluate {
        job_id: u64,
        invocable_name: String,
        input_data: String,
    }
}

#[cw_serde]
#[derive(QueryResponses)]
pub enum QueryMsg {}
