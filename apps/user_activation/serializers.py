from rest_framework import serializers


class UserActivationRegisterSerializer(serializers.Serializer):
    username = serializers.CharField(required=True)
    email = serializers.CharField(required=True)
    password = serializers.CharField(required=True)
    confirm_password = serializers.CharField(required=True)

    def validate(self, attrs):
        if attrs['password'] != attrs['confirm_password']:
            raise serializers.ValidationError('Passwords do not match')

        return attrs


class UserActivationResendActivationSerializer(serializers.Serializer):
    username = serializers.CharField(required=True)
