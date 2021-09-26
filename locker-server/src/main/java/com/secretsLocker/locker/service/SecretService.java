package com.secretsLocker.locker.service;

import com.secretsLocker.locker.dto.CreateSecretDto;
import com.secretsLocker.locker.dto.GetSecretDto;
import com.secretsLocker.locker.entity.Environment;
import com.secretsLocker.locker.entity.Repository;
import com.secretsLocker.locker.entity.Secret;
import com.secretsLocker.locker.exception.SecretException;
import com.secretsLocker.locker.repository.EnvironmentRepository;
import com.secretsLocker.locker.repository.SecretRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class SecretService {

    @Autowired
    EnvironmentRepository environmentRepository;

    @Autowired
    SecretRepository secretRepository;

    @Autowired
    EnvironmentService environmentService;

    @Autowired
    RepoService repoService;

    public Secret findByName(Environment env, String secretName) {
        Secret secret = null;
        for (Secret s : env.secrets) if (s.name.equals(secretName)) secret = s;
        return secret;
    }

    public Secret findByNameOrThrow(Environment env, String secretName) {
        Secret secret = this.findByName(env, secretName);
        if (secret == null) throw new SecretException.SecretDoesNotExist();
        return secret;
    }

    public void create(CreateSecretDto createSecretDto) {
        Repository repo = repoService.findByName(createSecretDto.repoName);
        Environment env = environmentService.findByName(repo, createSecretDto.envName);
        Secret secret = this.findByName(env, createSecretDto.secretName);

        if (secret != null) throw new SecretException.SecretAlreadyExists();

        Secret newSecret = new Secret();

        newSecret.name = createSecretDto.secretName;
        newSecret.value = createSecretDto.secretValue;

        env.secrets.add(newSecret);
        environmentRepository.save(env);
    }

    public void update(CreateSecretDto createSecretDto) {
        Repository repo = repoService.findByName(createSecretDto.repoName);
        Environment env = environmentService.findByName(repo, createSecretDto.envName);
        Secret secret = this.findByNameOrThrow(env, createSecretDto.secretName);

        secret.value = createSecretDto.secretValue;
        secretRepository.save(secret);
    }

    public String get(GetSecretDto getSecretDto) {
        Repository repo = repoService.findByName(getSecretDto.repoName);
        Environment env = environmentService.findByName(repo, getSecretDto.envName);
        Secret secret = this.findByNameOrThrow(env, getSecretDto.secretName);

        return secret.value;
    }
}
