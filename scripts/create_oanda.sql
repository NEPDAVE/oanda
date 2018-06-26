CREATE DATABASE oanda;

\c oanda

CREATE TABLE eur_usd_prices (
uuid serial PRIMARY KEY,
type text,
bid NUMERIC,
bid_liquidity NUMERIC,
ask NUMERIC,
ask_liquidity NUMERIC,
closeout_ask NUMERIC,
closeout_bid NUMERIC,
instrument text,
status text,
time timestamp
);

\q
