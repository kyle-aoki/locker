package com.secretsLocker.locker.service;

import com.secretsLocker.locker.dto.CreateSecretDto;
import com.secretsLocker.locker.dto.GetSecretDto;
import com.secretsLocker.locker.entity.Environment;
import com.secretsLocker.locker.entity.Repository;
import com.secretsLocker.locker.entity.Secret;
import com.secretsLocker.locker.exception.EnvironmentException;
import com.secretsLocker.locker.exception.RepoException;
import com.secretsLocker.locker.exception.SecretException;
import com.secretsLocker.locker.repository.EnvironmentRepository;
import com.secretsLocker.locker.repository.RepoRepository;
import com.secretsLocker.locker.repository.UserRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class SecretService {

    @Autowired
    RepoRepository repoRepository;

    public void create(CreateSecretDto createSecretDto) {
        Repository repo = repoRepository.findByName(createSecretDto.repoName);

        if (repo == null) throw new RepoException.RepoDoesNotExist();

        Environment env = null;
        for (Environment e : repo.environments) {
            if (e.name.equals(createSecretDto.envName)) env = e;
        }
        if (env == null) throw new EnvironmentException.EnvironmentNotFound();

        for (Secret s : env.secrets) {
            if (s.name.equals(createSecretDto.secretName)) throw new SecretException.SecretAlreadyExists();
        }

        Secret newSecret = new Secret();

        newSecret.name = createSecretDto.secretName;
        newSecret.value = createSecretDto.secretValue;

        env.secrets.add(newSecret);
        repoRepository.save(repo);
    }

    public String get(GetSecretDto getSecretDto) {
        Repository repo = repoRepository.findByName(getSecretDto.repoName);
        if (repo == null) throw new RepoException.RepoDoesNotExist();

        Environment env = null;
        for (Environment e : repo.environments) {
            if (e.name.equals(getSecretDto.envName)) env = e;
        }
        if (env == null) throw new EnvironmentException.EnvironmentNotFound();

        Secret secretToFind = null;
        for (Secret s : env.secrets) {
            if (s.name.equals(getSecretDto.secretName)) secretToFind = s;
        }
        if (secretToFind == null) throw new SecretException.SecretDoesNotExist();

        return secretToFind.value;
    }
}
