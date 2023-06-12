/* eslint-disable @typescript-eslint/no-unused-vars */
import { useQuery, type UseQueryOptions, useInfiniteQuery, type UseInfiniteQueryOptions } from "@tanstack/react-query";
import { useClient } from '../useClient';
import type { Ref } from 'vue'

export default function useDhpcUser() {
  const client = useClient();
  const QueryParams = ( options: any) => {
    const key = { type: 'QueryParams',  };    
    return useQuery([key], () => {
      return  client.DhpcUser.query.queryParams().then( res => res.data );
    }, options);
  }
  
  const QueryLinkedRequesters = (index: string,  options: any) => {
    const key = { type: 'QueryLinkedRequesters',  index };    
    return useQuery([key], () => {
      const { index } = key
      return  client.DhpcUser.query.queryLinkedRequesters(index).then( res => res.data );
    }, options);
  }
  
  const QueryLinkedRequestersAll = (query: any, options: any, perPage: number) => {
    const key = { type: 'QueryLinkedRequestersAll', query };    
    return useInfiniteQuery([key], ({pageParam = 1}: { pageParam?: number}) => {
      const {query } = key

      query['pagination.limit']=perPage;
      query['pagination.offset']= (pageParam-1)*perPage;
      query['pagination.count_total']= true;
      return  client.DhpcUser.query.queryLinkedRequestersAll(query ?? undefined).then( res => ({...res.data,pageParam}) );
    }, {...options,
      getNextPageParam: (lastPage, allPages) => { if ((lastPage.pagination?.total ?? 0) >((lastPage.pageParam ?? 0) * perPage)) {return lastPage.pageParam+1 } else {return undefined}},
      getPreviousPageParam: (firstPage, allPages) => { if (firstPage.pageParam==1) { return undefined } else { return firstPage.pageParam-1}}
    }
    );
  }
  
  const QuerySlashHistory = (index: string,  options: any) => {
    const key = { type: 'QuerySlashHistory',  index };    
    return useQuery([key], () => {
      const { index } = key
      return  client.DhpcUser.query.querySlashHistory(index).then( res => res.data );
    }, options);
  }
  
  const QuerySlashHistoryAll = (query: any, options: any, perPage: number) => {
    const key = { type: 'QuerySlashHistoryAll', query };    
    return useInfiniteQuery([key], ({pageParam = 1}: { pageParam?: number}) => {
      const {query } = key

      query['pagination.limit']=perPage;
      query['pagination.offset']= (pageParam-1)*perPage;
      query['pagination.count_total']= true;
      return  client.DhpcUser.query.querySlashHistoryAll(query ?? undefined).then( res => ({...res.data,pageParam}) );
    }, {...options,
      getNextPageParam: (lastPage, allPages) => { if ((lastPage.pagination?.total ?? 0) >((lastPage.pageParam ?? 0) * perPage)) {return lastPage.pageParam+1 } else {return undefined}},
      getPreviousPageParam: (firstPage, allPages) => { if (firstPage.pageParam==1) { return undefined } else { return firstPage.pageParam-1}}
    }
    );
  }
  
  const QueryUser = (account: string,  options: any) => {
    const key = { type: 'QueryUser',  account };    
    return useQuery([key], () => {
      const { account } = key
      return  client.DhpcUser.query.queryUser(account).then( res => res.data );
    }, options);
  }
  
  const QueryUserAll = (query: any, options: any, perPage: number) => {
    const key = { type: 'QueryUserAll', query };    
    return useInfiniteQuery([key], ({pageParam = 1}: { pageParam?: number}) => {
      const {query } = key

      query['pagination.limit']=perPage;
      query['pagination.offset']= (pageParam-1)*perPage;
      query['pagination.count_total']= true;
      return  client.DhpcUser.query.queryUserAll(query ?? undefined).then( res => ({...res.data,pageParam}) );
    }, {...options,
      getNextPageParam: (lastPage, allPages) => { if ((lastPage.pagination?.total ?? 0) >((lastPage.pageParam ?? 0) * perPage)) {return lastPage.pageParam+1 } else {return undefined}},
      getPreviousPageParam: (firstPage, allPages) => { if (firstPage.pageParam==1) { return undefined } else { return firstPage.pageParam-1}}
    }
    );
  }
  
  return {QueryParams,QueryLinkedRequesters,QueryLinkedRequestersAll,QuerySlashHistory,QuerySlashHistoryAll,QueryUser,QueryUserAll,
  }
}