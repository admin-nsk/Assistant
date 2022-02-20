from django import forms
from records.models import Record


class RecordForm(forms.ModelForm):

    class Meta:
        model = Record
        fields = (
            'title',
            'text',
            'date_of_creation',
            'date_of_diary',
            'tags'
        )