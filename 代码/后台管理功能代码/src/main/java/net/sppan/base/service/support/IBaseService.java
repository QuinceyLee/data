package net.sppan.base.service.support;

import org.springframework.data.domain.Example;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.data.domain.Sort;

import java.io.Serializable;
import java.util.List;

public interface IBaseService<T, ID extends Serializable> {
    T find(ID id);

    List<T> findAll();

    List<T> findList(ID[] ids);

    List<T> findList(Iterable<ID> ids);

    Page<T> findAll(Pageable pageable);

    Page<T> findAll(Example<T> spec, Pageable pageable);

    long count();

    long count(Example<T> spec);

    boolean exists(ID id);

    void save(T entity);

    T update(T entity);

    void delete(ID id);

    void deleteByIds(@SuppressWarnings("unchecked") ID... ids);

    void delete(T[] entitys);

    void delete(Iterable<T> entitys);

    void delete(T entity);

    List<T> findList(Example<T> spec, Sort sort);
}