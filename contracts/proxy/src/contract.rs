#[cfg(not(feature = "library"))]
use cosmwasm_std::entry_point;
use cosmwasm_std::{Binary, Deps, DepsMut, Env, MessageInfo, Response, StdResult};
// use cw2::set_contract_version;

use crate::error::ContractError;
use crate::msg::{ExecuteMsg, InstantiateMsg, QueryMsg, IbcExecuteMsg};
use cosmwasm_std::{to_json_binary, IbcMsg, IbcTimeout};

/*
// version info for migration info
const CONTRACT_NAME: &str = "crates.io:proxy";
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
    env: Env,
    _info: MessageInfo,
    msg: ExecuteMsg,
) -> Result<Response, ContractError> {
    match msg {
        ExecuteMsg::Evaluate {channel, job_id, invocable_name, input_data } => {
            Ok(Response::new()
                .add_attribute("method", "evaluate")
                .add_attribute("channel", channel.clone())
                .add_message(IbcMsg::SendPacket {
                    channel_id: channel,
                    data: to_json_binary(&IbcExecuteMsg::Evaluate {
                        job_id: job_id,
                        invocable_name: invocable_name,
                        input_data: input_data,
                    })?,
                    // default timeout of two minutes.
                    timeout: IbcTimeout::with_timestamp(env.block.time.plus_seconds(120)),
                }))
        }
    }
}

#[cfg_attr(not(feature = "library"), entry_point)]
pub fn query(_deps: Deps, _env: Env, _msg: QueryMsg) -> StdResult<Binary> {
    unimplemented!()
}

#[cfg(test)]
mod tests {}
