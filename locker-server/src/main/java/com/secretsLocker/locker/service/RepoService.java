package com.secretsLocker.locker.service;

import com.secretsLocker.locker.dto.ListRepoDto;
import com.secretsLocker.locker.dto.copy.CopyRepoDto;
import com.secretsLocker.locker.dto.delete.DeleteRepoDto;
import com.secretsLocker.locker.dto.path.RepoPath;
import com.secretsLocker.locker.dto.UpdateRepoDto;
import com.secretsLocker.locker.entity.Environment;
import com.secretsLocker.locker.entity.Repository;
import com.secretsLocker.locker.entity.Secret;
import com.secretsLocker.locker.entity.User;
import com.secretsLocker.locker.exception.Err;
import com.secretsLocker.locker.repository.EnvironmentRepository;
import com.secretsLocker.locker.repository.RepoRepository;
import com.secretsLocker.locker.repository.SecretRepository;
import com.secretsLocker.locker.repository.UserRepository;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import javax.persistence.EntityManager;
import java.util.ArrayList;
import java.util.List;

@Service
public class RepoService {

    Logger logger = LoggerFactory.getLogger(RepoService.class);

    @Autowired
    UserRepository userRepository;
    @Autowired
    SecretRepository secretRepository;
    @Autowired
    EnvironmentRepository environmentRepository;
    @Autowired
    RepoRepository repoRepository;

    public Repository findByName(String repoName) {
        return repoRepository.findByName(repoName);
    }

    public Repository findByNameOrThrow(String repoName) {
        Repository repo = repoRepository.findByName(repoName);
        if (repo == null) throw new Err("REPO_NOT_FOUND", "Repository not found.");
        return repo;
    }

    public void create(String username, RepoPath repoPath) {
        User user = userRepository.findByUsername(username);

        Repository potentialRepo = this.findByName(repoPath.repoName);
        if (potentialRepo != null) throw new Err("REPO_NAME_TAKEN", "Repository already exists.");

        Repository repository = new Repository();

        repository.setName(repoPath.repoName);
        repository.setOwner(user);

        repoRepository.save(repository);
    }

    public void rename(String username, UpdateRepoDto updateRepoDto) {
        Repository repo = this.findByNameOrThrow(updateRepoDto.repoName);

        repo.name = updateRepoDto.newRepoName;
        repoRepository.save(repo);
    }

    public List<String> listRepos(ListRepoDto listRepoDto) {
        int limit = listRepoDto.limit;
        int offset = listRepoDto.offset;
        long count = repoRepository.count();

        logger.info("limit=" + limit + " offset=" + offset);

        if (limit != -1 && offset == -1) return repoRepository.findAllNamesWithLimit(limit);
        if (limit == -1 && offset == -1) {
            if (count > 100) throw new Err("TOO_MANY_REPOS", "Would return too many repos.");
            return repoRepository.findAllNames();
        }

        return repoRepository.findAllWithLimitAndOffset(limit, offset);
    }

    public void delete(DeleteRepoDto deleteRepoDto) {
        Repository repo = this.findByNameOrThrow(deleteRepoDto.repoName);

        boolean hasEnvs = repo.environments.size() != 0;

        if (hasEnvs && !deleteRepoDto.force) {
            throw new Err("DEL_REPO_ERR", "This repository has existing environments. Use --force or delete them first.");
        }

        List<Secret> secretsToDelete = new ArrayList<>();
        List<Environment> environmentsToDelete = new ArrayList<>();

        for (Environment e : repo.environments) {
            secretsToDelete.addAll(e.secrets);
            e.secrets.clear();
            environmentsToDelete.add(e);
        }
        repo.environments.clear();

        secretRepository.deleteAll(secretsToDelete);
        environmentRepository.deleteAll(environmentsToDelete);
        repoRepository.delete(repo);
    }

    public void copy(CopyRepoDto copyRepoDto) {
        Repository repo = this.findByNameOrThrow(copyRepoDto.repoName);

        Repository repoClone = new Repository();
        repoClone.name = copyRepoDto.newRepoName;
        repoClone.owner = repo.owner;
        repoClone.members = new ArrayList<>();
        repoClone.environments = new ArrayList<>();

        List<Environment> envsToSave = new ArrayList<>();
        List<Secret> secretsToSave = new ArrayList<>();

        for (Environment env : repo.environments) {
            Environment envClone = new Environment();
            envClone.name = env.name;
            envClone.members = new ArrayList<>();

            envClone.secrets = new ArrayList<>();
            for (Secret secret : env.secrets) {
                Secret secretClone = new Secret();
                secretClone.name = secret.name;
                secretClone.value = secret.value;
                envClone.secrets.add(secretClone);
                secretsToSave.add(secretClone);
            }
            repoClone.environments.add(envClone);
            envsToSave.add(envClone);
        }

        secretRepository.saveAll(secretsToSave);
        environmentRepository.saveAll(envsToSave);
        repoRepository.save(repoClone);
    }
}





















