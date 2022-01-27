#!/usr/bin/env python3
import ccxt


class Bybit:
    """Bybitクラス
    """
    def __init__(self, apikey, secret, urls):
        """コンストラクタ

        Args:
            apikey (str): apiキー文字列
            secret (str): apiシークレットキー文字列
            urls (str): apiのURL文字列
        """
        self.bybit_api_inf              = {}
        self.bybit_api_inf["apikey"]    = apikey
        self.bybit_api_inf["secret"]    = secret
        self.bybit_api_inf["urls"]      = {"api":urls}
        self.bybit_api                  = ccxt.bybit(self.bybit_api_inf)
        
    def get_position(self):
        pass
        