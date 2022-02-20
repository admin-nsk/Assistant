from django import forms


class RecordForm(forms.ModelForm):

    class Meta:
        fields = (
            'title',
            'text',
            'date_of_creation',
            'date_of_diary',
            'tags'
        )