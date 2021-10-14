package com.secretsLocker.locker.dto.copy;

import com.secretsLocker.locker.dto.path.RepoEnvPath;

public class CopyEnv extends RepoEnvPath {
    public String targetRepoName;
    public String targetEnvName;

    public String path() {
        return this.repoName + "/" + this.envName;
    }
    public String targetPath() {
        return this.targetRepoName + "/" + this.targetEnvName;
    }
}
