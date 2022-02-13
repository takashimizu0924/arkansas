#!/usr/bin/env python3
# -*- coding: utf-8 -*-
import ccxt
from asset.base_asset import BaseAsset

class Bybit(BaseAsset):
    """Bybitクラス
    """
    def setup(self, api_info_dict=None):
        """ api初期化処理
        """
        if api_info_dict:
            return ccxt.bybit(api_info_dict)
        return ccxt.bybit()

    def fetch_ticker(self,symbol):
        """ティッカー情報を取得する
            NOTE:BybitのPublicAPI

        Args:
            symbol (str): 取得するティッカーシンボルの指定
        Returns:
            str: ticker情報 / None
        """
        ticker = None
        if not symbol == "":
            ticker = self.api.fetch_ticker(symbol)
        return ticker
    
    def fetch_my_trades(self, symbol):
        trades = self.api.fetch_my_trades(symbol)
        return trades