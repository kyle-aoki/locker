package com.secretsLocker.locker.service;

import com.secretsLocker.locker.dto.*;
import com.secretsLocker.locker.dto.delete.DeleteEnvDto;
import com.secretsLocker.locker.dto.diff.MissingRequest;
import com.secretsLocker.locker.dto.path.RepoEnvPath;
import com.secretsLocker.locker.entity.Environment;
import com.secretsLocker.locker.entity.Repository;
import com.secretsLocker.locker.entity.Secret;
import com.secretsLocker.locker.exception.Err;
import com.secretsLocker.locker.repository.EnvironmentRepository;
import com.secretsLocker.locker.repository.RepoRepository;
import com.secretsLocker.locker.repository.SecretRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.ArrayList;
import java.util.List;
import java.util.stream.Collectors;

@Service
public class EnvironmentService {

    @Autowired
    RepoRepository repoRepository;

    @Autowired
    RepoService repoService;

    @Autowired
    SecretService secretService;

    @Autowired
    EnvironmentRepository environmentRepository;

    @Autowired
    SecretRepository secretRepository;

    public Environment findByNameOrThrow(Repository repo, String envName) {
        Environment env = null;
        for (Environment e : repo.environments) {
            if (e.name.equals(envName)) env = e;
        }
        if (env == null) throw new Err("ENV_NOT_FOUND", "Environment not found.");
        return env;
    }

    public void create(RepoEnvPath repoEnvPath) {
        Repository repo = repoService.findByNameOrThrow(repoEnvPath.repoName);

        String newEnvName = repoEnvPath.envName;

        for (Environment env : repo.environments) {
            if (env.name.equals(newEnvName)) throw new Err("ENV_NAME_TAKEN", "Environment name taken.");
        }
        Environment env = new Environment(newEnvName);

        repo.environments.add(env);

        environmentRepository.save(env);
        repoRepository.save(repo);
    }

    public List<KeyValue> get(RepoEnvPath repoEnvPath) {
        Repository repo = repoService.findByNameOrThrow(repoEnvPath.repoName);
        Environment env = this.findByNameOrThrow(repo, repoEnvPath.envName);

        List<KeyValue> keyValues = new ArrayList<>();
        for (Secret s : env.secrets) keyValues.add(new KeyValue(s.name, s.value));

        return keyValues;
    }

    public List<String> missing(MissingRequest missingRequest) {
        Repository repo = repoService.findByNameOrThrow(missingRequest.repoName);
        Environment givenEnv = this.findByNameOrThrow(repo, missingRequest.envName);
        Environment targetEnv = this.findByNameOrThrow(repo, missingRequest.targetEnvName);

        List<Secret> givenEnvSecrets = new ArrayList<>(givenEnv.secrets);
        List<Secret> targetEnvSecrets = new ArrayList<>(targetEnv.secrets);

        targetEnvSecrets.removeAll(givenEnvSecrets);

        return targetEnvSecrets.stream()
                .map(secret -> secret.name)
                .collect(Collectors.toList());
    }

    public void delete(DeleteEnvDto deleteEnvDto) {
        Repository repo = repoService.findByNameOrThrow(deleteEnvDto.repoName);
        Environment env = this.findByNameOrThrow(repo, deleteEnvDto.envName);

        boolean hasSecrets = env.secrets.size() != 0;

        if (hasSecrets && !deleteEnvDto.force) {
            throw new Err("NOT_EMPTY", "This environment is not empty. Use --force or delete secrets first.");
        }

        for (Secret s : env.secrets) secretRepository.delete(s);
        env.secrets.clear();

        environmentRepository.delete(env);
        repo.environments.remove(env);

        repoRepository.save(repo);
    }
}





















