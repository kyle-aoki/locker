package com.secretsLocker.locker.cryptography;

import javax.persistence.*;
import java.util.Date;
import java.util.Random;

public class SessionToken {
    private static final String LowercaseLetters = "abcdefghijklmnopqrstuvwxyz";
    private static final String UppercaseLetters = "ABCDEFGHIJKLMNOPQRSTUVEXYZ";
    private static final String Numbers = "1234567890";
    private static final char[] CharPool = (LowercaseLetters + UppercaseLetters + Numbers).toCharArray();

    private static final int SessionTokenLength = 40;

    // 1000 milliseconds  *  60 seconds  *  60 minutes  *  24 hours  =  86_400_000 milliseconds
    //         second            minute          hour          day                      day
    private static final Long SessionTokenValidityDuration = 1000L * 60L * 60L * 24L;
    private static final Long StringComparisonMinTime = 30L;

    @Column
    private final String sessionTokenKey;

    @Temporal(TemporalType.TIMESTAMP)
    private Date creationTime;

    public SessionToken() {
        this.sessionTokenKey = this.generateRandomAlphaNumericString();
        this.creationTime = new Date();
    }

    private String generateRandomAlphaNumericString() {
        Random rand = new Random();
        StringBuilder sb = new StringBuilder();

        for (int i = 0; i < SessionTokenLength; i += 1) {
            int x = rand.nextInt(CharPool.length);
            sb.append(CharPool[x]);
        }

        return sb.toString();
    }

    public Date getExpireDate() {
        Long creationTimeMilliseconds = this.creationTime.toInstant().toEpochMilli();
        return new Date(creationTimeMilliseconds + SessionTokenValidityDuration);
    }

    public String getSessionTokenKey() {
        return sessionTokenKey;
    }

    // Protects against timing attacks
    public static Boolean timeSafeStringComparison(String sessionTokenInDB, String sessionTokenStringFromRequest) {
        long t1 = System.currentTimeMillis();
        boolean isValidSessionToken = sessionTokenInDB.equals(sessionTokenStringFromRequest);
        long delta = System.currentTimeMillis() - t1;
        long timeToWait = StringComparisonMinTime - delta;
        try {
            Thread.sleep(timeToWait);
        } catch (InterruptedException ignored) {
        }
        return isValidSessionToken;
    }

    public Date getCreationTime() {
        return creationTime;
    }

    public void setCreationTime(Date creationTime) {
        this.creationTime = creationTime;
    }
}
