#!/usr/bin/env python3

from as_def_data import As_MarketData
from config.config import Parser
from asset.bybit import Bybit

class AS_TradeMarket:
    """ 市場クラス
    """
    def __init__(self, auto_mode=False, test_mode=False, market_type="bybit"):
        """コンストラクタ

        Args:
            auto_mode (bool, optional): 自動取引モード設定 Defaults to False.
            test_mode (bool, optional): テストモード設定 Defaults to False.
            market_type (str, optional): 市場タイプ名 Defaults to "bybit".
        """
        self.is_auto = auto_mode
        self.is_test = test_mode
        self.is_reload = False
        self.market = None
        self.market_data = As_MarketData()
        self.config_parser = None
        self.api_key = ""
        self.api_secret = ""
        if self.check_market_type(market_type):
            self.market_type = market_type
        else:
            self.market_type = "bybit"
        self.__preproc()

    def __preproc(self):
        api_key = ""
        api_secret = ""
        url = ""
        # 自動取引モード設定
        if self.get_auto_mode():
            config_file_path = "/config/config.ini"
            config_parser = Parser(config_file_path)
            api_key = config_parser.get_sectionsValue_fromKey(self.market_type, "api_key")
            api_secret = config_parser.get_sectionsValue_fromKey(self.market_type, "api_secret")
            # テストモード設定
            if self.get_test_mode():
                url = config_parser.get_sectionsValue_fromKey("bybit", "test_base_url")
        # 市場別処理
        if self.market_type == "bybit":
            self.market = Bybit(api_key, api_secret, url)
        else:
            raise Exception(f"Market type error. market_type:{self.market_type}")
    
    def __reload(self):
        if self.is_reload:
            self.__preproc()

    def check_market_type(self, market_type):
        return self.market_data.is_exist_into_table(market_type)

    def set_auto_mode(self, auto_mode):
        self.is_auto = auto_mode

    def set_test_mode(self, test_mode):
        self.is_test = test_mode
    
    def get_auto_mode(self):
        return self.is_auto
    
    def get_test_mode(self):
        return self.is_test
    
    def fetch_ticker(self, symbol):
        return self.market.fetch_ticker(symbol)
            
    # def get_ticker(self, symbol_type,symbol_id):
    #     symbol = BybitSymbol.get_name_from_id(symbol_type, symbol_id)
    #     return self.bybit.fetch_ticker(symbol)        