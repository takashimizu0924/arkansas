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
        
    def fetch_ticker(self,symbol):
        """ティッカー情報を取得する

        Args:
            symbol (str): 取得するティッカーシンボルの指定
        Returns:
            ticker (object) / None
        """
        if not symbol:
            return None
        ticker = self.bybit_api.fetch_ticker(symbol)
        return ticker