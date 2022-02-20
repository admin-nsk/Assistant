from selenium import webdriver
import unittest


class NewVisitorTest(unittest.TestCase):
    """тест страницы пользователя"""

    def setUp(self) -> None:
        self.browser = webdriver.Firefox()

    def tearDown(self) -> None:
        self.browser.quit()

    def test_index_page_for_user(self):
        """тест: открывается главная страница дневника"""
        # Саша входит на страницу Личного дневника
        # Он видит что заголовок и шапка сайта говорят о личном дневнике
        self.browser.get('http://localhost:8000')
        assert 'Личный дневник' in self.browser.title

    def test_can_create_resord_in_diary(self):
        """тест: можно создать запись в дневнике"""
        # Ему предлагается создать новую запись
        # Он набирает в теме записи "Моя первая запись"
        # В поле текста вводит "Это моя первая запись в этом дневнике"
        # Выбирает дату записи "01.01.2022"
        # Нажимает создать
        # И видит появившуюся запись под формой создания записи
        pass


if __name__ == '__main__':
    unittest.main(warnings='ignore')
