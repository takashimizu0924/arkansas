#!/usr/bin/env python3
# -*- coding: utf-8 -*-
import ccxt


class BybitSymbol:
    def __init__(self):
        self.__symbol_table = {
                                "USDT":[
                                    "BIT/USDT",
                                    "BTC/USDT",
                                    "ETH/USDT",
                                    "MANA/USDT",
                                    "SAND/USDT",
                                    "SHIB1000/USDT",
                                    "ADA/USDT",
                                    "BNB/USDT",
                                    "XRP/USDT",
                                    "SOL/USDT",
                                    "DOT/USDT",
                                    "DOGE/USDT",
                                    "UNI/USDT",
                                    "LUNA/USDT",
                                    "AVAX/USDT",
                                    "LINK/USDT",
                                    "LTC/USDT",
                                    "ALGO/USDT",
                                    "BCH/USDT",
                                    "ATOM/USDT",
                                    "MATIC/USDT",
                                    "FIL/USDT",
                                    "ICP/USDT",
                                    "ETC/USDT",
                                    "XLM/USDT",
                                    "VET/USDT",
                                    "AXS/USDT",
                                    "TRX/USDT",
                                    "FTT/USDT",
                                    "XTZ/USDT",
                                    "THETA/USDT",
                                    "HBAR/USDT",
                                    "EOS/USDT",
                                    "AAVE/USDT",
                                    "NEAR/USDT",
                                    "FTM/USDT",
                                    "KSM/USDT",
                                    "OMG/USDT",
                                    "IOTX/USDT",
                                    "DASH/USDT",
                                    "COMP/USDT",
                                    "ONE/USDT",
                                    "CHZ/USDT",
                                    "LRC/USDT",
                                    "ENJ/USDT",
                                    "XEM/USDT",
                                    "SUSHI/USDT",
                                    "DYDX/USDT",
                                    "SRM/USDT",
                                    "CRV/USDT",
                                    "IOST/USDT",
                                    "CELR/USDT",
                                    "CHR/USDT",
                                    "WOO/USDT",
                                    "ALICE/USDT",
                                    "ENS/USDT",
                                    "C98/USDT",
                                    "GALA/USDT",
                                    "KEEP/USDT",
                                    "SLP/USDT",
                                    "ANT/USDT",
                                    "ROSE/USDT",
                                    "CROU/USDT",
                                    "TFI/USDT",
                                    "KNC/USDT",
                                    "SXP/USDT",
                                    "AR/USDT",
                                    "STORJ/USDT",
                                    "SPELL/USDT",
                                    "BSV/USDT",
                                    "CREAM/USDT",
                                    "COTI/USDT",
                                    "1INCHI/USDT",
                                    "TLM/USDT",
                                    "EGLD/USDT",
                                    "PEOPLE/USDT",
                                    "FLOW/USDT",
                                    "COMP/USDT",
                                    "CELOU/USDT",
                                    "BAT/USDT",
                                    "ZEN/USDT",
                                    "STX/USDT",
                                    "SNX/USDT",
                                    "XMR/USDT",
                                    "SRM/USDT",
                                    "DUSK/USDT",
                                    "WAVES/USDT",
                                    "YGG/USDT",
                                    "DENT/USDT",
                                    "IOTA/USDT",
                                    "RVN/USDT",
                                    "ANKR/USDT",
                                    "KAVA/USDT",
                                    "RNDRU/USDT",
                                    "RUNE/USDT",
                                    "ILVU/USDT",
                                    "QTUM/USDT",
                                    "CTK/USDT",
                                    "ZEC/USDT",
                                    "BTT/USDT",
                                    "LOOKS/USDT",
                                    "KLAY/USDT",
                                    "SFP/USDT",
                                    "GTC/USDT",
                                    "REQ/USDT",
                                    "GRT/USDT",
                                    "MASK/USDT",
                                    "AUDIO/USDT",
                                    "LPY/USDT",
                                    "IMX/USDT",
                                    "HNT/USDT",
                                    "CVC/USDT",
                                    "JASMY/USDT",
                                    "BICO/USDT",
                                    "RSR/USDT",
                                    "REN/USDT",
                                    "LIT/USDT",
                                    "SCU/USDT",
                                    ]
                                }
        def get_symbol_max_id(self, symbol_type):
            return len(self.__symbol_table[symbol_type])
            
        def get_name_from_id(self, symbol_type, symbol_id):
            """ID番号から通貨種別名を取得する

            Args:
                symbol_type (str): 通貨種別
                symbol_id (int): [description]

            Returns:
                [type]: [description]
            """
            if not self.__symbol_table[symbol_type] or not self.__symbol_table[symbol_type][symbol_id]:
                return ""
            return self.__symbol_table[symbol_type][symbol_id]
            


class Bybit:
    """Bybitクラス
    """
    def __init__(self, apikey="", secret="", urls=""):
        """コンストラクタ
           NOTE: bybitAPIを使用する事前準備
        Args:
            apikey (str, optional): apiキー文字列, Defaults to "".
            secret (str, optional): apiシークレットキー文字列, Defaults to "".
            urls (str, optional): apiのURL文字列, Defaults to "".
        """
        if not apikey == "" and not secret == "":
            self.bybit_api_inf              = {}
            self.bybit_api_inf["apiKey"]    = apikey
            self.bybit_api_inf["secret"]    = secret
            self.bybit_api_inf["urls"]      = {"api":urls}  # 不要情報？
            self.bybit_api                  = ccxt.bybit(self.bybit_api_inf)
        else:
            self.bybit_api                  = ccxt.bybit()
        
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
            ticker = self.bybit_api.fetch_ticker(symbol)
        return ticker
    
    # def fetch_my_trades(self, symbol):
    #     trades = self.bybit_api.fetch_my_trades(symbol)
    #     return trades