#!/usr/bin/env python3
import configparser

class Parser:
    """パーサークラス
    """
    def __init__(self, config_file_path):
        """コンストラクタ

        Args:
            config_file_path (str): configファイルへのパス
        """
        self.config = configparser.ConfigParser()
        self.config_file_path = config_file_path
        self._pre_read_config()
        
    def _pre_read_config(self):
        """ 事前にconfigファイルを読み込む
        """
        if self.config_file_path == "" or not self.config_file_path in ".ini":
            self.config_file_path = "config.ini"
        self.config.read(self.config_file_path)

    def get_sectionsValue_fromKey(self, section_name, key):
        """対象セクションのキーから値を取得する

        Args:
            section_name (str): 対象セクション名
            key (str): 対象キー名

        Returns:
            str: 設定データ文字列
        """
        sections_value = None
        sections_info = self.config.sections()
        if  section_name in sections_info:
            sections_value = self.config[section_name][key]
        return sections_value
        

