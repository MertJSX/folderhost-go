import { useState, useEffect, useRef, useCallback } from 'react';
import Cookies from 'js-cookie';
const API_BASE_URL = import.meta.env.VITE_API_BASE_URL;

const useWebSocket = (path: string, shouldConnect: boolean) => {
  const wsRef = useRef<WebSocket | null>(null);
  const [isConnected, setIsConnected] = useState<Boolean>(false);
  const [connectionError, setConnectionError] = useState<Boolean>(false);
  const isConnectedRef = useRef<Boolean>(false);
  const [messages, setMessages] = useState<Array<string>>([]);
  const reconnectTimeoutRef = useRef<number | null>(null);

  const previousPathRef = useRef<string>(path);

  const connect = useCallback(() => {
    if (wsRef.current) {
      wsRef.current.close();
      wsRef.current = null;
    }

    const token = Cookies.get("token");
    if (!token) {
      console.error('Token not found');
      return;
    }

    const wsUrl = `${API_BASE_URL}/ws/${encodeURIComponent(path)}?token=${encodeURIComponent(token)}`;

    const websocket = new WebSocket(wsUrl);

    websocket.onopen = () => {
      console.log('WebSocket has been connected!');
      setIsConnected(true);
      isConnectedRef.current = true;
      wsRef.current = websocket;
      setConnectionError(false);

      if (reconnectTimeoutRef.current) {
        clearTimeout(reconnectTimeoutRef.current);
        reconnectTimeoutRef.current = null;
      }
    };

    websocket.onmessage = (event) => {
      setMessages(prev => [...prev, event.data]);
    };

    websocket.onclose = (event) => {
      console.log('WebSocket connection has been closed', event.code, event.reason);
      setIsConnected(false);
      isConnectedRef.current = false;
      wsRef.current = null;

      if (event.code !== 1000 && shouldConnect && !reconnectTimeoutRef.current) {
        setConnectionError(true);
        reconnectTimeoutRef.current = setTimeout(() => {
          console.log('Trying to connect again...');
          connect();
        }, 5000);
      }
    };

    websocket.onerror = (error) => {
      console.error('WebSocket error:', error);
      setConnectionError(true);
    };
  }, [path, shouldConnect]);

  const disconnect = useCallback(() => {
    if (wsRef.current) {
      wsRef.current.close(1000, "Manual disconnect");
      wsRef.current = null;
    }
    if (reconnectTimeoutRef.current) {
      clearTimeout(reconnectTimeoutRef.current);
      reconnectTimeoutRef.current = null;
    }
    setIsConnected(false);
    isConnectedRef.current = false;
  }, []);

  const sendMessage = useCallback((message: string) => {
    if (wsRef.current && wsRef.current.readyState === WebSocket.OPEN) {
      wsRef.current.send(message);
    } else {
      console.warn("WebSocket not ready, cannot send:", message);
    }
  }, []);

  useEffect(() => {
    if (shouldConnect) {
      connect();
    } else {
      disconnect();
    }

    return () => {
      disconnect();
    };
  }, [shouldConnect]);

  useEffect(() => {
    if (shouldConnect && isConnectedRef.current && previousPathRef.current !== path) {
      previousPathRef.current = path;
      const delta = JSON.stringify({
        type: 'change-path',
        path: path?.slice(1)
      });
      sendMessage(delta)
    }
  }, [path, shouldConnect]);

  return {
    isConnected,
    isConnectedRef,
    messages,
    sendMessage,
    connectionError,
    connect,
    disconnect
  };
};

export default useWebSocket;