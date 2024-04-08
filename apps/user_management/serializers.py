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


class UserManagementResetPasswordSerializer(serializers.Serializer):
    username = serializers.CharField(required=True)


class UserManagementChangeEmailSerializer(serializers.Serializer):
    old_email = serializers.EmailField(required=True)
    new_email = serializers.EmailField(required=True)

    def validate(self, attrs):
        if attrs['new_email'] == attrs['old_email']:
            raise serializers.ValidationError('Email must be different from old email')

        return attrs


class UserManagementChangePasswordSerializer(serializers.Serializer):
    old_password = serializers.CharField(required=True)
    new_password = serializers.CharField(required=True)
    confirm_new_password = serializers.CharField(required=True)

    def validate(self, attrs):
        if attrs['new_password'] != attrs['confirm_new_password']:
            raise serializers.ValidationError('Passwords do not match')

        if attrs['new_password'] == attrs['old_password']:
            raise serializers.ValidationError('Password must be different from old password')

        return attrs
