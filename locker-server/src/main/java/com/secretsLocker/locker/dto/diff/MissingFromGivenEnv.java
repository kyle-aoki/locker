package com.secretsLocker.locker.dto.diff;

import java.util.ArrayList;
import java.util.List;

public class MissingFromGivenEnv {
    public List<String> missingSecrets;

    public MissingFromGivenEnv() {
        this.missingSecrets = new ArrayList<>();
    }
}
