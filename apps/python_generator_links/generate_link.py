import json
import sys

from sys import argv


class GeneratorLinks:
    def __init__(self):
        self.rabbit_url = {}
        self.db_pg_url = {}
        self.db_msql_url = {}
        self.cabinet_url = {}
        self.crm2_url = {}
        self.mail_hog_url = {}
        self.old_services_url = {}
        self.new_services_url = {}

    def generate(self, request=None):
        self.rabbit_url = {"host": "db.test" + request + ".test:15672",
                           "login": "guest",
                           "password": "guest",
                           }
        self.db_pg_url = {"host": "db.test" + request + ".test:5432",
                          "db": "fbs",
                          "login": "fbs",
                          "password": "f65_p455",
                          }
        self.db_msql_url = {"host": "db.test" + request + ".test:3306",
                            "db": "cabinet, mt4real{NS}, NS - номер сервера(1-10), mt4demo, mt5to4real1, mt5to4demo",
                            "login": "fbs",
                            "password": "fL5485lGn8qF",
                            }
        self.cabinet_url = {"host": "https://cabinet" + request + ".fbs"}
        self.crm2_url = {"host": "http://crm2.cabinet" + request + ".fbs",
                         "login": "admin@fbs.com",
                         "password": "superfbs",
                         }
        self.mail_hog_url = {"host": "https://cabinet" + request + ".fbs:8025"}
        self.old_services_url = {"host": "apps-01.test-" + request + ".cabinet.office.int.fbs.dev"}
        self.new_services_url = {"host": "app-01.test-" + request + ".new-cabinet.office.fbs.dev"}

        return [self.rabbit_url,
                self.db_pg_url,
                self.db_msql_url,
                self.cabinet_url,
                self.crm2_url,
                self.mail_hog_url,
                self.old_services_url,
                self.new_services_url]


if __name__ == "__main__":
    if len(argv) == 2:
        Generator = GeneratorLinks()
        links = {"links": Generator.generate(argv[1])}
        json_links = json.dumps(links)
        print(json_links)

    else:
        print('incorrect len of args')
        sys.exit(1)
