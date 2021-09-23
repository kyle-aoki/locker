package com.secretsLocker.locker.cryptography;

import org.apache.commons.codec.digest.DigestUtils;
import org.springframework.beans.factory.annotation.Value;

public class LockerEncryption {

    @Value("${locker.encryption.key}")
    private static String lockerEncryptionKey;

    public static String encryptPassword(String plaintextPassword) {
        return DigestUtils.sha256Hex(plaintextPassword + lockerEncryptionKey);
    }
}
