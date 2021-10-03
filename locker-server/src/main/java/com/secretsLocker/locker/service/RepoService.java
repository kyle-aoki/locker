package com.secretsLocker.locker.service;

import com.secretsLocker.locker.dto.ListRepoDto;
import com.secretsLocker.locker.dto.path.RepoPath;
import com.secretsLocker.locker.dto.UpdateRepoDto;
import com.secretsLocker.locker.entity.Repository;
import com.secretsLocker.locker.entity.User;
import com.secretsLocker.locker.exception.Err;
import com.secretsLocker.locker.repository.RepoRepository;
import com.secretsLocker.locker.repository.UserRepository;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class RepoService {

    Logger logger = LoggerFactory.getLogger(RepoService.class);

    @Autowired
    UserRepository userRepository;

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

    public void update(String username, UpdateRepoDto updateRepoDto) {
        Repository repo = this.findByNameOrThrow(updateRepoDto.repoName);

        if (!repo.owner.username.equals(username)) {
            throw new Err("INSUFF_PERMS.", "Insufficient permissions to change repo name. Only repo owner may perform this action.");
        }

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
}
