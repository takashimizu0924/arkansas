#!/usr/bin/env python3

class As_MarketData:
    """ 市場種別クラス
    """
    def __init__(self):
        """ コンストラクタ
            NOTE:マーケット情報を設定する
        """
        self.__market_table = {
            "market-type":{
                "bybit":{
                    "symbol-type":{
                        "USDT":[
                            "BIT/USDT", "BTC/USDT", "ETH/USDT", "MANA/USDT", "SAND/USDT",
                            "SHIB1000/USDT", "ADA/USDT", "BNB/USDT", "XRP/USDT", "SOL/USDT",
                            "DOT/USDT", "DOGE/USDT", "UNI/USDT", "LUNA/USDT", "AVAX/USDT",
                            "LINK/USDT", "LTC/USDT", "ALGO/USDT", "BCH/USDT", "ATOM/USDT",
                            "MATIC/USDT", "FIL/USDT", "ICP/USDT", "ETC/USDT", "XLM/USDT",
                            "VET/USDT", "AXS/USDT", "TRX/USDT", "FTT/USDT", "XTZ/USDT",
                            "THETA/USDT", "HBAR/USDT", "EOS/USDT", "AAVE/USDT", "NEAR/USDT",
                            "FTM/USDT", "KSM/USDT", "OMG/USDT", "IOTX/USDT", "DASH/USDT",
                            "COMP/USDT", "ONE/USDT", "CHZ/USDT", "LRC/USDT", "ENJ/USDT",
                            "XEM/USDT", "SUSHI/USDT", "DYDX/USDT", "SRM/USDT", "CRV/USDT",
                            "IOST/USDT", "CELR/USDT", "CHR/USDT", "WOO/USDT", "ALICE/USDT",
                            "ENS/USDT", "C98/USDT", "GALA/USDT", "KEEP/USDT", "SLP/USDT",
                            "ANT/USDT", "ROSE/USDT", "CROU/USDT", "TFI/USDT", "KNC/USDT",
                            "SXP/USDT", "AR/USDT", "STORJ/USDT", "SPELL/USDT", "BSV/USDT",
                            "CREAM/USDT", "COTI/USDT", "1INCHI/USDT", "TLM/USDT", "EGLD/USDT",
                            "PEOPLE/USDT", "FLOW/USDT", "COMP/USDT", "CELOU/USDT", "BAT/USDT",
                            "ZEN/USDT", "STX/USDT", "SNX/USDT", "XMR/USDT", "SRM/USDT",
                            "DUSK/USDT", "WAVES/USDT", "YGG/USDT", "DENT/USDT", "IOTA/USDT",
                            "RVN/USDT", "ANKR/USDT", "KAVA/USDT", "RNDRU/USDT", "RUNE/USDT",
                            "ILVU/USDT", "QTUM/USDT", "CTK/USDT", "ZEC/USDT", "BTT/USDT",
                            "LOOKS/USDT", "KLAY/USDT", "SFP/USDT", "GTC/USDT", "REQ/USDT",
                            "GRT/USDT", "MASK/USDT", "AUDIO/USDT", "LPY/USDT", "IMX/USDT",
                            "HNT/USDT", "CVC/USDT", "JASMY/USDT", "BICO/USDT", "RSR/USDT",
                            "REN/USDT", "LIT/USDT", "SCU/USDT"
                        ],
                    }
                }
            }
        }
    
    def is_exist_into_table(self, target_type_str):
        """ 種別情報の存在確認をする

        Args:
            target_type_str (str): 種別情報文字列(市場種別、シンボル種別[通貨種別名]、シンボル名[通貨名])
        
        Returns:
            bool: True:存在する/False:存在しない
        """
        is_exist = False
        for market_type_key, market_type_val in self.__market_table["market-type"].items():
            for symbol_type_key, symbol_type_val in market_type_val["symbol-type"].items():
                for symbol_name in symbol_type_val:
                    if symbol_name == target_type_str:
                        is_exist = True
                        break
                if symbol_type_key == target_type_str:
                    is_exist = True
                    break
            if market_type_key == target_type_str:
                is_exist = True
                break
        return is_exist
    
    def get_length_from_symbol_name_table(self, market_type, market_symbol_type):
        """ シンボル名のテーブル長を返す

        Args:
            market_type (str): 市場種別
            market_symbol_type (str): シンボル種別[通貨種別名]

        Returns:
            int: シンボル名テーブル長
        """
        length = 0
        try:
            length = len(self.__market_table["market-type"][market_type]["symbol-type"][market_symbol_type])
        except:
            length = -1
        return length

    def get_symbol_name_from_id(self, market_type, market_symbol_type, symbol_id):
        """ IDからシンボル名を取得する

        Args:
            market_type (str): 市場種別
            market_symbol_type (str): シンボル種別[通貨種別名]
            symbol_id (int): シンボルID

        Returns:
            str: シンボル名
        """
        symbol_name_table_length = self.get_length_from_symbol_name_table(market_type, market_symbol_type)
        symbol_name = ""
        if 0 < symbol_name_table_length:
            try:
                symbol_name = self.__market_table["market-type"][market_type]["symbol-type"][market_symbol_type][symbol_id]
            except:
                symbol_name = ""
        return symbol_name