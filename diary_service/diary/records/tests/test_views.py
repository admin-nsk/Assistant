from django.test import TestCase
from django.urls import resolve
from records.views import index


class HomePageTest(TestCase):
    """тест главной страницы"""

    def test_root_url_resolve_to_home_page(self):
        """тест: по корневому url открыватся представление главной страницы"""
        found = resolve('/')
        self.assertEqual(found.func, index)