package com.secretsLocker.locker.service;

import com.secretsLocker.locker.dto.CreateRepoDto;
import com.secretsLocker.locker.dto.ListRepoDto;
import com.secretsLocker.locker.dto.UpdateRepoDto;
import com.secretsLocker.locker.entity.Repository;
import com.secretsLocker.locker.entity.User;
import com.secretsLocker.locker.exception.RepoException;
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
        Repository repo = repoRepository.findByName(repoName);
        if (repo == null) throw new RepoException.RepoDoesNotExist();
        return repo;
    }

    public void create(String username, CreateRepoDto createRepoDto) {
        User user = userRepository.findByUsername(username);

        Repository potentialRepo = repoRepository.findByName(createRepoDto.repoName);
        if (potentialRepo != null) throw new RepoException.RepoAlreadyExists();

        Repository repository = new Repository();

        repository.setName(createRepoDto.repoName);
        repository.setOwner(user);

        repoRepository.save(repository);
    }

    public void update(String username, UpdateRepoDto updateRepoDto) {
        Repository repo = this.findByName(updateRepoDto.repoName);

        if (!repo.owner.username.equals(username)) {
            throw new RepoException.NotRepoOwner("You are unauthorized to change the repo's name.");
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
            if (count > 100) throw new RepoException.TooManyRepos();
            return repoRepository.findAllNames();
        }

        return repoRepository.findAllWithLimitAndOffset(limit, offset);
    }
}
