from datetime import datetime

from django.db import models


class Record(models.Model):
    """Модель заметки"""
    title = models.TextField(verbose_name='Тема')
    text = models.TextField(verbose_name='Текст записи')
    date_of_creation = models.DateTimeField(verbose_name='Дата создания', default=datetime.now())
    date_of_diary = models.DateTimeField(verbose_name='Дата дневника')
    tags = models.ManyToManyField('Tag', related_name='records', verbose_name='Тэги')

    def __str__(self):
        return f'{self.title} {self.date_of_diary}'


class Tag(models.Model):
    """Модель тега"""
    name = models.CharField(max_length=100)

    def __str__(self):
        return self.name
