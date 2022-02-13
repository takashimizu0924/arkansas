#!/usr/bin/env python3
# -*- coding: utf-8 -*-

class BaseAsset:
    """ベースクラス
    """
    def __init__(self, apikey="", secret="", urls=""):
        """コンストラクタ
        Args:
            apikey (str, optional): apiキー文字列, Defaults to "".
            secret (str, optional): apiシークレットキー文字列, Defaults to "".
            urls (str, optional): apiのURL文字列, Defaults to "".
        """
        self.api_inf = {}
        self.api = None
        if apikey == "" or secret == "":
            if not urls == "" and "test" in urls:
                self.api_inf["urls"] = {"api":urls}
            else:
                self.api = self.setup()
        else:
            self.api_inf["apiKey"] = apikey
            self.api_inf["secret"] = secret
        self.api = self.setup(self.api_inf)
        
    def setup(self, object=None):
        """ api初期化処理継承用
        """
        pass

    def fetch_ticker(self,symbol):
        """ティッカー情報取得継承用
        Args:
            symbol (str): 取得するティッカーシンボルの指定
        Returns:
            str: ticker情報 / None
        """
        pass