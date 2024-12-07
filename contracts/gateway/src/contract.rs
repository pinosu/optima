#[cfg(not(feature = "library"))]
use cosmwasm_std::entry_point;
use cosmwasm_std::{Binary, Deps, DepsMut, Env, MessageInfo, Response, CosmosMsg, StdResult};
// use cw2::set_contract_version;

use crate::error::ContractError;
use crate::msg::{ExecuteMsg, InstantiateMsg, QueryMsg, OptimaMsg};

/*
// version info for migration info
const CONTRACT_NAME: &str = "crates.io:gateway";
const CONTRACT_VERSION: &str = env!("CARGO_PKG_VERSION");
*/

#[cfg_attr(not(feature = "library"), entry_point)]
pub fn instantiate(
    _deps: DepsMut,
    _env: Env,
    _info: MessageInfo,
    _msg: InstantiateMsg,
) -> Result<Response, ContractError> {
    Ok(Response::default())
}

#[cfg_attr(not(feature = "library"), entry_point)]
pub fn execute(
    _deps: DepsMut,
    _env: Env,
    _info: MessageInfo,
    msg: ExecuteMsg,
) -> Result<Response<OptimaMsg>, ContractError> {
    match msg {
        ExecuteMsg::Evaluate { job_id, invocable_name, input_data } => {
            evaluate(job_id, invocable_name, input_data)
        }
    }
}

#[cfg_attr(not(feature = "library"), entry_point)]
pub fn query(_deps: Deps, _env: Env, _msg: QueryMsg) -> StdResult<Binary> {
    unimplemented!()
}

fn evaluate(
    job_id: u64,
    invocable_name: String,
    input_data: String,
) -> Result<Response<OptimaMsg>, ContractError> {
    let custom_msg = OptimaMsg::Invocable {
        job_id,
        invocable_name,
        input_data,
    };
    Ok(Response::new().add_message(CosmosMsg::Custom(custom_msg)))
}
