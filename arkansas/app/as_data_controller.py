#!/usr/bin/env python3

from arkansas.app.config import Parser
from arkansas.app.asset.bybit import Bybit, BybitSymbol

class AS_TradeMarket:
    """ 市場クラス
    """
    def __init__(self, auto_mode=False, test_mode=False):
        if auto_mode:
            self.config_file_path = "/config/config.ini"
            self.config_parser = Parser(self.config_file_path)
            self.config_bybit_api_key = self.config_parser.get_sectionsValue_fromKey("bybit", "api_key")
            self.config_bybit_api_secret = self.config_parser.get_sectionsValue_fromKey("bybit", "api_secret")
            self.config_bybit_urls = ""
            if test_mode:
                self.config_bybit_urls = self.config_parser.get_sectionsValue_fromKey("bybit", "test_base_url")
            else:
                self.config_bybit_urls = self.config_parser.get_sectionsValue_fromKey("bybit", "base_url")
            self.bybit = Bybit(self.config_bybit_api_key, self.config_bybit_api_secret, self.config_bybit_urls)
        else:
            self.bybit = Bybit()
            
    def get_ticker(self, symbol_type,symbol_id):
        symbol = BybitSymbol.get_name_from_id(symbol_type, symbol_id)
        return self.bybit.fetch_ticker(symbol)
        