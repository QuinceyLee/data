package net.sppan.base.service.support.impl;

import java.io.Serializable;
import java.util.Arrays;
import java.util.List;

import javax.transaction.Transactional;

import net.sppan.base.dao.support.IBaseDao;
import net.sppan.base.entity.support.BaseEntity;
import net.sppan.base.service.support.IBaseService;

import org.springframework.data.domain.Example;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.data.domain.Sort;

@Transactional
public abstract class BaseServiceImpl<T extends BaseEntity, ID extends Serializable> implements IBaseService<T, ID> {

    public abstract IBaseDao<T, ID> getBaseDao();

    @Override
    public T find(ID id) {
        return getBaseDao().findById(id).get();
    }

    @Override	
    public List<T> findAll() {
        return getBaseDao().findAll();
    }

    @Override
    public List<T> findList(ID[] ids) {
        List<ID> idList = Arrays.asList(ids);
        return getBaseDao().findAllById(idList);
    }

    @Override
    public List<T> findList(Example<T> spec, Sort sort) {
        return getBaseDao().findAll(spec, sort);
//                findAll(spec, sort);
    }

    @Override
    public Page<T> findAll(Pageable pageable) {
        return getBaseDao().findAll(pageable);
    }

    @Override
    public long count() {
        return getBaseDao().count();
    }

    @Override
    public long count(Example<T> spec) {
        return getBaseDao().count(spec);
    }

    @Override
    public boolean exists(ID id) {
        return getBaseDao().existsById(id);
    }

    @Override
    public void save(T entity) {
        getBaseDao().save(entity);
    }

    public void save(Iterable<T> entitys) {
        getBaseDao().saveAll(entitys);
    }

    @Override
    public T update(T entity) {
        return getBaseDao().saveAndFlush(entity);
    }

    @Override
    public void delete(ID id) {
        getBaseDao().deleteById(id);
    }

    @Override
    public void deleteByIds(@SuppressWarnings("unchecked") ID... ids) {
        if (ids != null) {
            for (int i = 0; i < ids.length; i++) {
                ID id = ids[i];
                this.delete(id);
            }
        }
    }

    @Override
    public void delete(T[] entitys) {
        List<T> tList = Arrays.asList(entitys);
        getBaseDao().deleteInBatch(tList);
    }

    @Override
    public void delete(Iterable<T> entitys) {
        getBaseDao().deleteInBatch(entitys);
    }

    @Override
    public void delete(T entity) {
        getBaseDao().delete(entity);
    }

    @Override
    public List<T> findList(Iterable<ID> ids) {
        return getBaseDao().findAllById(ids);
    }

    @Override
    public Page<T> findAll(Example<T> spec, Pageable pageable) {
        // TODO Auto-generated method stub
        return getBaseDao().findAll(spec, pageable);
    }

}
