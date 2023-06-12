/* eslint-disable @typescript-eslint/no-unused-vars */
import { useQuery, type UseQueryOptions, useInfiniteQuery, type UseInfiniteQueryOptions } from "@tanstack/react-query";
import { useClient } from '../useClient';
import type { Ref } from 'vue'

export default function useDhpcRequest() {
  const client = useClient();
  const QueryParams = ( options: any) => {
    const key = { type: 'QueryParams',  };    
    return useQuery([key], () => {
      return  client.DhpcRequest.query.queryParams().then( res => res.data );
    }, options);
  }
  
  const QueryAllowedOracles = (id: string,  options: any) => {
    const key = { type: 'QueryAllowedOracles',  id };    
    return useQuery([key], () => {
      const { id } = key
      return  client.DhpcRequest.query.queryAllowedOracles(id).then( res => res.data );
    }, options);
  }
  
  const QueryAllowedOraclesAll = (query: any, options: any, perPage: number) => {
    const key = { type: 'QueryAllowedOraclesAll', query };    
    return useInfiniteQuery([key], ({pageParam = 1}: { pageParam?: number}) => {
      const {query } = key

      query['pagination.limit']=perPage;
      query['pagination.offset']= (pageParam-1)*perPage;
      query['pagination.count_total']= true;
      return  client.DhpcRequest.query.queryAllowedOraclesAll(query ?? undefined).then( res => ({...res.data,pageParam}) );
    }, {...options,
      getNextPageParam: (lastPage, allPages) => { if ((lastPage.pagination?.total ?? 0) >((lastPage.pageParam ?? 0) * perPage)) {return lastPage.pageParam+1 } else {return undefined}},
      getPreviousPageParam: (firstPage, allPages) => { if (firstPage.pageParam==1) { return undefined } else { return firstPage.pageParam-1}}
    }
    );
  }
  
  const QueryMinerResponse = (uUID: string,  options: any) => {
    const key = { type: 'QueryMinerResponse',  uUID };    
    return useQuery([key], () => {
      const { uUID } = key
      return  client.DhpcRequest.query.queryMinerResponse(uUID).then( res => res.data );
    }, options);
  }
  
  const QueryMinerResponseAll = (query: any, options: any, perPage: number) => {
    const key = { type: 'QueryMinerResponseAll', query };    
    return useInfiniteQuery([key], ({pageParam = 1}: { pageParam?: number}) => {
      const {query } = key

      query['pagination.limit']=perPage;
      query['pagination.offset']= (pageParam-1)*perPage;
      query['pagination.count_total']= true;
      return  client.DhpcRequest.query.queryMinerResponseAll(query ?? undefined).then( res => ({...res.data,pageParam}) );
    }, {...options,
      getNextPageParam: (lastPage, allPages) => { if ((lastPage.pagination?.total ?? 0) >((lastPage.pageParam ?? 0) * perPage)) {return lastPage.pageParam+1 } else {return undefined}},
      getPreviousPageParam: (firstPage, allPages) => { if (firstPage.pageParam==1) { return undefined } else { return firstPage.pageParam-1}}
    }
    );
  }
  
  const QueryRequestRecord = (uUID: string,  options: any) => {
    const key = { type: 'QueryRequestRecord',  uUID };    
    return useQuery([key], () => {
      const { uUID } = key
      return  client.DhpcRequest.query.queryRequestRecord(uUID).then( res => res.data );
    }, options);
  }
  
  const QueryRequestRecordAll = (query: any, options: any, perPage: number) => {
    const key = { type: 'QueryRequestRecordAll', query };    
    return useInfiniteQuery([key], ({pageParam = 1}: { pageParam?: number}) => {
      const {query } = key

      query['pagination.limit']=perPage;
      query['pagination.offset']= (pageParam-1)*perPage;
      query['pagination.count_total']= true;
      return  client.DhpcRequest.query.queryRequestRecordAll(query ?? undefined).then( res => ({...res.data,pageParam}) );
    }, {...options,
      getNextPageParam: (lastPage, allPages) => { if ((lastPage.pagination?.total ?? 0) >((lastPage.pageParam ?? 0) * perPage)) {return lastPage.pageParam+1 } else {return undefined}},
      getPreviousPageParam: (firstPage, allPages) => { if (firstPage.pageParam==1) { return undefined } else { return firstPage.pageParam-1}}
    }
    );
  }
  
  const QueryRequestRecordAllMinerPending = (query: any, options: any) => {
    const key = { type: 'QueryRequestRecordAllMinerPending', query };    
    return useQuery([key], () => {
      const {query } = key
      return  client.DhpcRequest.query.queryRequestRecordAllMinerPending(query ?? undefined).then( res => res.data );
    }, options);
  }
  
  const QueryRequestRecordAllAnswerPending = (query: any, options: any) => {
    const key = { type: 'QueryRequestRecordAllAnswerPending', query };    
    return useQuery([key], () => {
      const {query } = key
      return  client.DhpcRequest.query.queryRequestRecordAllAnswerPending(query ?? undefined).then( res => res.data );
    }, options);
  }
  
  const QueryTreasury = ( options: any) => {
    const key = { type: 'QueryTreasury',  };    
    return useQuery([key], () => {
      return  client.DhpcRequest.query.queryTreasury().then( res => res.data );
    }, options);
  }
  
  return {QueryParams,QueryAllowedOracles,QueryAllowedOraclesAll,QueryMinerResponse,QueryMinerResponseAll,QueryRequestRecord,QueryRequestRecordAll,QueryRequestRecordAllMinerPending,QueryRequestRecordAllAnswerPending,QueryTreasury,
  }
}