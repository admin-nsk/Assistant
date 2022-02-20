from django.urls import path
from records import views


urlpatterns = [
    path('', views.index, name='index'),
]