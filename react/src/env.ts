const apiURL = import.meta.env.VITE_API_COSMOS ?? "http://localhost:1317";
const rpcURL = import.meta.env.VITE_WS_TENDERMINT ?? "http://localhost:26657";
const prefix = "dhpc";

export const env = {
  apiURL,
  rpcURL,
  prefix,
};