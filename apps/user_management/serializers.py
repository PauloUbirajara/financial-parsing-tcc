from rest_framework import serializers


class UserManagementRegisterSerializer(serializers.Serializer):
    username = serializers.CharField(required=True)
    email = serializers.CharField(required=True)
    password = serializers.CharField(required=True)
    confirm_password = serializers.CharField(required=True)

    def validate(self, attrs):
        if attrs['password'] != attrs['confirm_password']:
            raise serializers.ValidationError('Passwords do not match')

        return attrs


class UserManagementResendActivationSerializer(serializers.Serializer):
    username = serializers.CharField(required=True)
