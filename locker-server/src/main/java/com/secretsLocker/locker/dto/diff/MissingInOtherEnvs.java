package com.secretsLocker.locker.dto.diff;

import java.util.List;

public class MissingInOtherEnvs {
    public String envName;
    public List<String> missingSecrets;

    public MissingInOtherEnvs (String envName, List<String> missingSecrets) {
        this.envName = envName;
        this.missingSecrets = missingSecrets;
    }
}
